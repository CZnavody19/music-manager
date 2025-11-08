package setup

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/CZnavody19/music-manager/src/config"
	"github.com/CZnavody19/music-manager/src/db"
	configStore "github.com/CZnavody19/music-manager/src/db/config"
	musicbrainzStore "github.com/CZnavody19/music-manager/src/db/musicbrainz"
	youtubeStore "github.com/CZnavody19/music-manager/src/db/youtube"
	"github.com/CZnavody19/music-manager/src/graph"
	"github.com/CZnavody19/music-manager/src/graph/generated"
	"github.com/CZnavody19/music-manager/src/http"
	"github.com/CZnavody19/music-manager/src/internal/auth"
	"github.com/CZnavody19/music-manager/src/internal/discord"
	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/plex"
	"github.com/CZnavody19/music-manager/src/internal/youtube"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewResolver(dbConn *sql.DB, mqConn *amqp.Connection, config config.Config) (*graph.Resolver, error) {
	dbMapper := db.NewMapper()

	configStore := configStore.NewConfigStore(dbConn, dbMapper)
	youtubeStore := youtubeStore.NewYouTubeStore(dbConn, dbMapper)
	musibcrainzStore := musicbrainzStore.NewMusicbrainzStore(dbConn, dbMapper)

	mb, err := musicbrainz.NewMusicBrainz(musibcrainzStore)
	if err != nil {
		return nil, err
	}

	auth, err := auth.NewAuth(configStore, config.Server.TokenCheckEnable)
	if err != nil {
		return nil, err
	}

	yt, err := youtube.NewYouTube(configStore, youtubeStore, mb)
	if err != nil {
		return nil, err
	}

	dsc, err := discord.NewDiscord(configStore)
	if err != nil {
		return nil, err
	}

	plx, err := plex.NewPlex(configStore)
	if err != nil {
		return nil, err
	}

	httpHandler := http.NewHttpHandler(configStore)

	graphMapper := graph.NewMapper()
	graphInputMapper := graph.NewInputMapper()

	return &graph.Resolver{
		Mapper:      graphMapper,
		InputMapper: graphInputMapper,
		MusicBrainz: mb,
		Auth:        auth,
		YouTube:     yt,
		Discord:     dsc,
		Plex:        plx,
		HttpHandler: httpHandler,
		ConfigStore: configStore,
	}, nil
}

func SetupDirectives(config *generated.Config, directives *graph.Directives) {
	config.Directives.Auth = directives.Auth

	config.Directives.DiscordEnabled = directives.DiscordEnabled
	config.Directives.PlexEnabled = directives.PlexEnabled
	config.Directives.YoutubeEnabled = directives.YoutubeEnabled
}

func SetupPresenters(srv *handler.Server) {
}
