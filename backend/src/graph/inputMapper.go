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

func (im *InputMapper) MapDiscordConfigInput(input model.DiscordConfigInput) *domain.DiscordConfig {
	return &domain.DiscordConfig{
		WebhookURL: input.WebhookURL,
	}
}

func (im *InputMapper) MapPlexConfigInput(input model.PlexConfigInput) *domain.PlexConfig {
	return &domain.PlexConfig{
		Protocol:  input.Protocol,
		Host:      input.Host,
		Port:      input.Port,
		Token:     input.Token,
		LibraryID: input.LibraryID,
	}
}
