package db

import (
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/domain"
)

type Mapper struct {
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) MapYouTubeConfig(input *model.Youtube) *domain.YouTubeConfig {
	return &domain.YouTubeConfig{
		OAuth: input.OAuth,
		Token: input.Token,
	}
}

func (m *Mapper) MapDiscordConfig(input *model.Discord) *domain.DiscordConfig {
	return &domain.DiscordConfig{
		WebhookURL: input.WebhookURL,
	}
}
