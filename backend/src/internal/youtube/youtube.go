package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/db/youtube"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/websockets"
	"github.com/CZnavody19/music-manager/src/mq"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/sosodev/duration"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	youtubeApi "google.golang.org/api/youtube/v3"
)

type YouTube struct {
	enabled     bool
	configStore *config.ConfigStore
	config      *domain.YouTubeConfig
	ytStore     *youtube.YouTubeStore
	yt          *youtubeApi.Service
	musicBrainz *musicbrainz.MusicBrainz
	websockets  *websockets.Websockets
	mq          *mq.MessageQueue
}

func getYtService(ctx context.Context, cfg *domain.YouTubeConfig) (*youtubeApi.Service, error) {
	if cfg == nil {
		return nil, nil
	}

	config, err := google.ConfigFromJSON(cfg.OAuth, "https://www.googleapis.com/auth/youtube.readonly")
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.Unmarshal(cfg.Token, t)
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, t)

	yt, err := youtubeApi.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return yt, nil
}

func NewYouTube(cs *config.ConfigStore, yts *youtube.YouTubeStore, mb *musicbrainz.MusicBrainz, ws *websockets.Websockets, mq *mq.MessageQueue) (*YouTube, error) {
	ctx := context.Background()

	config, err := cs.GetYoutubeConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	service, err := getYtService(ctx, config)
	if err != nil {
		return nil, err
	}

	enabled := false
	if service != nil && config.Enabled {
		enabled = true
	}

	return &YouTube{
		enabled:     enabled,
		configStore: cs,
		config:      config,
		ytStore:     yts,
		yt:          service,
		musicBrainz: mb,
		websockets:  ws,
		mq:          mq,
	}, nil
}

func (yt *YouTube) IsEnabled() bool {
	return yt.enabled
}

func (yt *YouTube) Enable(ctx context.Context) error {
	config, err := yt.configStore.GetYoutubeConfig(ctx)
	if err != nil {
		return err
	}

	service, err := getYtService(ctx, config)
	if err != nil {
		return err
	}

	err = yt.configStore.SetYoutubeEnabled(ctx, true)
	if err != nil {
		return err
	}

	yt.yt = service
	yt.config = config
	yt.enabled = true

	return nil
}

func (yt *YouTube) Disable(ctx context.Context) error {
	err := yt.configStore.SetYoutubeEnabled(ctx, false)
	if err != nil {
		return err
	}

	yt.yt = nil
	yt.config = nil
	yt.enabled = false

	return nil
}

func (yt *YouTube) RefreshPlaylist(ctx context.Context) error {
	if !yt.enabled {
		return nil
	}

	start := time.Now()
	yt.websockets.SendTask(&model.Task{
		Title:     "Refreshing YouTube playlist",
		StartedAt: start,
		Ended:     false,
	})

	zap.S().Info("Refreshing YouTube playlist")

	pageToken, err := yt.ytStore.GetLatestPageToken(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return err
	}

	res, err := yt.yt.PlaylistItems.List([]string{"snippet"}).Context(ctx).PlaylistId(yt.config.PlaylistID).MaxResults(50).PageToken(pageToken).Do()
	if err != nil {
		return err
	}

	token := res.NextPageToken
	videos := make(map[string]*domain.YouTubeVideo, len(res.Items))
	ids := make([]string, 0, len(res.Items))

	if token == "" {
		token = pageToken
	}

	for _, item := range res.Items {
		ids = append(ids, item.Snippet.ResourceId.VideoId)

		videos[item.Snippet.ResourceId.VideoId] = &domain.YouTubeVideo{
			VideoID:       item.Snippet.ResourceId.VideoId,
			Title:         item.Snippet.Title,
			ChannelTitle:  item.Snippet.VideoOwnerChannelTitle,
			ThumbnailURL:  &item.Snippet.Thumbnails.Default.Url,
			Duration:      nil,
			Position:      item.Snippet.Position,
			NextPageToken: token,
		}
	}

	vidRes, err := yt.yt.Videos.List([]string{"contentDetails"}).Context(ctx).Id(strings.Join(ids, ",")).MaxResults(50).Do()
	if err != nil {
		return err
	}

	videoArr := make([]*domain.YouTubeVideo, 0, len(videos))

	for _, item := range vidRes.Items {
		dur, err := duration.Parse(item.ContentDetails.Duration)
		if err != nil {
			return err
		}

		duration := int64(dur.ToTimeDuration().Seconds())

		origItem := videos[item.Id]
		origItem.Duration = &duration
		videoArr = append(videoArr, origItem)
	}

	err = yt.ytStore.StoreVideos(ctx, videoArr)
	if err != nil {
		return err
	}

	for _, video := range videoArr {
		yt.musicBrainz.SearchQueue <- IdentificationRequest{
			Video:   video,
			YtStore: yt.ytStore,
			Mq:      yt.mq,
		}
	}

	yt.websockets.SendTask(&model.Task{
		Title:     "Refreshed YouTube playlist",
		StartedAt: start,
		Ended:     true,
	})

	zap.S().Info("YouTube playlist refreshed successfully")

	return err
}

func (yt *YouTube) GetVideos(ctx context.Context) ([]*domain.YouTubeVideo, error) {
	if !yt.enabled {
		return nil, fmt.Errorf("YouTube integration is not enabled")
	}

	return yt.ytStore.GetVideos(ctx, false)
}

func (yt *YouTube) GetVideoByID(ctx context.Context, id string) (*domain.YouTubeVideo, error) {
	if !yt.enabled {
		return nil, fmt.Errorf("YouTube integration is not enabled")
	}

	return yt.ytStore.GetVideoByID(ctx, id)
}

func (yt *YouTube) MatchVideo(ctx context.Context, videoID string, trackID uuid.UUID) error {
	if !yt.enabled {
		return fmt.Errorf("YouTube integration is not enabled")
	}

	return yt.ytStore.LinkTrack(ctx, videoID, trackID)
}
