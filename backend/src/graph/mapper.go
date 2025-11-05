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
