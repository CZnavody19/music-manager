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

func (im *Mapper) MapYoutubeVideo(input *domain.YouTubeVideo) *model.YouTubeVideo {
	if input == nil {
		return nil
	}

	return &model.YouTubeVideo{
		ID:           input.VideoID,
		Title:        input.Title,
		ChannelTitle: input.ChannelTitle,
		ThumbnailURL: *input.ThumbnailURL,
		Duration:     *input.Duration,
		Position:     input.Position,
	}
}

func (im *Mapper) MapYoutubeVideos(inputs []*domain.YouTubeVideo) []*model.YouTubeVideo {
	var outputs []*model.YouTubeVideo
	for _, input := range inputs {
		outputs = append(outputs, im.MapYoutubeVideo(input))
	}
	return outputs
}
