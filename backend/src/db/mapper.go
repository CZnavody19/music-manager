package db

import (
	cfgModel "github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/model"
	"github.com/CZnavody19/music-manager/src/domain"
)

type Mapper struct {
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapYouTubeConfig(input *cfgModel.Youtube) *domain.YouTubeConfig {
	return &domain.YouTubeConfig{
		Enabled:    input.Enabled,
		OAuth:      input.OAuth,
		Token:      input.Token,
		PlaylistID: input.PlaylistID,
	}
}

func (m *Mapper) MapDiscordConfig(input *cfgModel.Discord) *domain.DiscordConfig {
	return &domain.DiscordConfig{
		Enabled:    input.Enabled,
		WebhookURL: input.WebhookURL,
	}
}

func (m *Mapper) MapPlexConfig(input *cfgModel.Plex) *domain.PlexConfig {
	return &domain.PlexConfig{
		Enabled:   input.Enabled,
		Protocol:  input.Protocol,
		Host:      input.Host,
		Port:      int64(input.Port),
		Token:     input.Token,
		LibraryID: int64(input.LibraryID),
	}
}

func (m *Mapper) MapGeneralConfig(input *cfgModel.General) *domain.GeneralConfig {
	return &domain.GeneralConfig{
		DownloadPath: input.DownloadPath,
		TempPath:     input.TempPath,
	}
}

func (m *Mapper) MapAuthConfig(input *cfgModel.Auth) *domain.AuthConfig {
	return &domain.AuthConfig{
		Username:     input.Username,
		PasswordHash: input.PasswordHash,
	}
}

func (m *Mapper) MapYoutubeVideos(input []*model.Youtube) []*domain.YouTubeVideo {
	var out []*domain.YouTubeVideo

	for _, v := range input {
		out = append(out, m.MapYoutubeVideo(v))
	}

	return out
}

func (m *Mapper) MapYoutubeVideo(input *model.Youtube) *domain.YouTubeVideo {
	var duration *int64
	if input.Duration != nil {
		d := int64(*input.Duration)
		duration = &d
	}

	return &domain.YouTubeVideo{
		VideoID:       input.VideoID,
		Title:         input.Title,
		ChannelTitle:  input.ChannelTitle,
		ThumbnailURL:  input.ThumbnailURL,
		Duration:      duration,
		Position:      int64(input.Position),
		NextPageToken: input.NextPageToken,
		TrackID:       input.TrackID,
	}
}

func (m *Mapper) MapPlexTrack(input *model.Plex) *domain.PlexTrack {
	if input == nil {
		return nil
	}

	return &domain.PlexTrack{
		ID:       int64(input.ID),
		Title:    input.Title,
		Artist:   input.Artist,
		Duration: int64(input.Duration),
		Mbid:     input.Mbid,
	}
}

func (m *Mapper) MapPlexTracks(input []*model.Plex) []*domain.PlexTrack {
	var out []*domain.PlexTrack

	for _, t := range input {
		out = append(out, m.MapPlexTrack(t))
	}

	return out
}

func (m *Mapper) MapTrackWithISRCs(input TrackWithISRCs) *domain.Track {
	var isrcs []string
	for _, isrc := range input.ISRCs {
		isrcs = append(isrcs, isrc.Isrc)
	}

	return &domain.Track{
		ID:     input.ID,
		Title:  input.Title,
		Artist: input.Artist,
		Length: int64(input.Length),
		ISRCs:  isrcs,
	}
}

func (m *Mapper) MapTracksWithISRCs(input []TrackWithISRCs) []*domain.Track {
	var out []*domain.Track

	for _, t := range input {
		out = append(out, m.MapTrackWithISRCs(t))
	}

	return out
}
