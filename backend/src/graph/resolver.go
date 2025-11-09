package graph

import (
	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/http"
	"github.com/CZnavody19/music-manager/src/internal/auth"
	"github.com/CZnavody19/music-manager/src/internal/discord"
	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/plex"
	"github.com/CZnavody19/music-manager/src/internal/websockets"
	"github.com/CZnavody19/music-manager/src/internal/youtube"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Mapper      *Mapper
	InputMapper *InputMapper

	Auth        *auth.Auth
	MusicBrainz *musicbrainz.MusicBrainz
	YouTube     *youtube.YouTube
	Discord     *discord.Discord
	Plex        *plex.Plex
	Websockets  *websockets.Websockets

	HttpHandler *http.HttpHandler

	ConfigStore *config.ConfigStore
}
