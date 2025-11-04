package graph

import (
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
)

type InputMapper struct {
}

func NewInputMapper() *InputMapper {
	return &InputMapper{}
}

func (im *InputMapper) MapDiscordConfig(input model.DiscordConfig) *domain.DiscordConfig {
	return &domain.DiscordConfig{
		WebhookURL: input.WebhookURL,
	}
}
