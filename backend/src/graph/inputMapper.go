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

func (im *InputMapper) MapLoginInput(input model.LoginInput) *domain.Credentials {
	return &domain.Credentials{
		Username: input.Username,
		Password: input.Password,
	}
}

func (im *InputMapper) MapYoutubeConfigInput(input model.YoutubeConfigInput) *domain.YouTubeConfig {
	return &domain.YouTubeConfig{
		PlaylistID: input.PlaylistID,
	}
}
