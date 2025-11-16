package graph

import (
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
)

type Mapper struct {
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (im *Mapper) MapGeneralConfig(input *domain.GeneralConfig) *model.GeneralConfig {
	return &model.GeneralConfig{
		DownloadPath: input.DownloadPath,
		TempPath:     input.TempPath,
	}
}

func (im *Mapper) MapDiscordConfig(input *domain.DiscordConfig) *model.DiscordConfig {
	return &model.DiscordConfig{
		WebhookURL: input.WebhookURL,
	}
}

func (im *Mapper) MapPlexConfig(input *domain.PlexConfig) *model.PlexConfig {
	return &model.PlexConfig{
		Protocol:  input.Protocol,
		Host:      input.Host,
		Port:      input.Port,
		Token:     input.Token,
		LibraryID: input.LibraryID,
	}
}

func (im *Mapper) MapYoutubeConfig(input *domain.YouTubeConfig) *model.YoutubeConfig {
	return &model.YoutubeConfig{
		PlaylistID: input.PlaylistID,
	}
}

func (im *Mapper) MapTidalConfig(input *domain.TidalConfig) *model.TidalConfig {
	return &model.TidalConfig{
		AuthTokenType:    input.AuthTokenType,
		AuthAccessToken:  input.AuthAccessToken,
		AuthRefreshToken: input.AuthRefreshToken,
		AuthExpiresAt:    input.AuthExpiresAt,
		AuthClientID:     input.AuthClientID,
		AuthClientSecret: input.AuthClientSecret,
		DownloadTimeout:  input.DownloadTimeout,
		DownloadRetries:  input.DownloadRetries,
		DownloadThreads:  input.DownloadThreads,
		AudioQuality:     input.AudioQuality,
	}
}

func (im *Mapper) MapYoutubeVideo(input *domain.YouTubeVideo) *model.YouTubeVideo {
	if input == nil {
		return nil
	}

	linked := false
	if input.TrackID != nil {
		linked = true
	}

	return &model.YouTubeVideo{
		ID:           input.VideoID,
		Title:        input.Title,
		ChannelTitle: input.ChannelTitle,
		ThumbnailURL: *input.ThumbnailURL,
		Duration:     *input.Duration,
		Position:     input.Position,
		Linked:       linked,
	}
}

func (im *Mapper) MapYoutubeVideos(inputs []*domain.YouTubeVideo) []*model.YouTubeVideo {
	var outputs []*model.YouTubeVideo
	for _, input := range inputs {
		outputs = append(outputs, im.MapYoutubeVideo(input))
	}
	return outputs
}

func (im *Mapper) MapTrack(input *domain.Track) *model.Track {
	if input == nil {
		return nil
	}

	return &model.Track{
		ID:            input.ID,
		Title:         input.Title,
		Artist:        input.Artist,
		Length:        input.Length,
		Isrcs:         input.ISRCs,
		LinkedYoutube: input.LinkedYoutube,
		LinkedPlex:    input.LinkedPlex,
	}
}

func (im *Mapper) MapTracks(inputs []*domain.Track) []*model.Track {
	var outputs []*model.Track
	for _, input := range inputs {
		outputs = append(outputs, im.MapTrack(input))
	}
	return outputs
}

func (im *Mapper) MapPlexTrack(input *domain.PlexTrack) *model.PlexTrack {
	if input == nil {
		return nil
	}

	return &model.PlexTrack{
		ID:       input.ID,
		Title:    input.Title,
		Artist:   input.Artist,
		Duration: input.Duration,
		Mbid:     input.Mbid,
		TrackID:  input.TrackID,
	}
}

func (im *Mapper) MapPlexTracks(inputs []*domain.PlexTrack) []*model.PlexTrack {
	var outputs []*model.PlexTrack
	for _, input := range inputs {
		outputs = append(outputs, im.MapPlexTrack(input))
	}
	return outputs
}
