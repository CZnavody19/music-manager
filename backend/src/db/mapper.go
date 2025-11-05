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
		Enabled: input.Enabled,
		OAuth:   input.OAuth,
		Token:   input.Token,
	}
}

func (m *Mapper) MapDiscordConfig(input *model.Discord) *domain.DiscordConfig {
	return &domain.DiscordConfig{
		Enabled:    input.Enabled,
		WebhookURL: input.WebhookURL,
	}
}

func (m *Mapper) MapPlexConfig(input *model.Plex) *domain.PlexConfig {
	return &domain.PlexConfig{
		Enabled:   input.Enabled,
		Protocol:  input.Protocol,
		Host:      input.Host,
		Port:      int64(input.Port),
		Token:     input.Token,
		LibraryID: int64(input.LibraryID),
	}
}

func (m *Mapper) MapGeneralConfig(input *model.General) *domain.GeneralConfig {
	return &domain.GeneralConfig{
		DownloadPath: input.DownloadPath,
		TempPath:     input.TempPath,
	}
}

func (m *Mapper) MapAuthConfig(input *model.Auth) *domain.AuthConfig {
	return &domain.AuthConfig{
		Username:     input.Username,
		PasswordHash: input.PasswordHash,
	}
}
