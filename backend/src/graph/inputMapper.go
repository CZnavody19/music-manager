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

func (im *InputMapper) MapGeneralConfigInput(input model.GeneralConfigInput) *domain.GeneralConfig {
	return &domain.GeneralConfig{
		DownloadPath: input.DownloadPath,
		TempPath:     input.TempPath,
	}
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

func (im *InputMapper) MapTidalConfigInput(input model.TidalConfigInput) *domain.TidalConfig {
	return &domain.TidalConfig{
		AuthTokenType:        input.AuthTokenType,
		AuthAccessToken:      input.AuthAccessToken,
		AuthRefreshToken:     input.AuthRefreshToken,
		AuthExpiresAt:        input.AuthExpiresAt,
		AuthClientID:         input.AuthClientID,
		AuthClientSecret:     input.AuthClientSecret,
		DownloadTimeout:      input.DownloadTimeout,
		DownloadRetries:      input.DownloadRetries,
		DownloadThreads:      input.DownloadThreads,
		AudioQuality:         input.AudioQuality,
		FilePermissions:      input.FilePermissions,
		DirectoryPermissions: input.DirectoryPermissions,
		Owner:                input.Owner,
		Group:                input.Group,
	}
}
