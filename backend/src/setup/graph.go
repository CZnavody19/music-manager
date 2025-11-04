package setup

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/CZnavody19/music-manager/src/config"
	"github.com/CZnavody19/music-manager/src/db"
	configStore "github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/graph"
	"github.com/CZnavody19/music-manager/src/graph/generated"
	"github.com/CZnavody19/music-manager/src/http"
	"github.com/CZnavody19/music-manager/src/internal/discord"
	"github.com/CZnavody19/music-manager/src/internal/youtube"
)

func NewResolver(dbConn *sql.DB, config config.Config) (*graph.Resolver, error) {
	dbMapper := db.NewMapper()

	configStore := configStore.NewConfigStore(dbConn, dbMapper)

	yt, err := youtube.NewYouTube(configStore)
	if err != nil {
		return nil, err
	}

	dsc, err := discord.NewDiscord(configStore)
	if err != nil {
		return nil, err
	}

	httpHandler := http.NewHttpHandler(configStore)

	graphInputMapper := graph.NewInputMapper()

	return &graph.Resolver{
		InputMapper: graphInputMapper,
		YouTube:     yt,
		Discord:     dsc,
		HttpHandler: httpHandler,
		ConfigStore: configStore,
	}, nil
}

func SetupDirectives(config *generated.Config, directives *graph.Directives) {
	config.Directives.DiscordEnabled = directives.DiscordEnabled
}

func SetupPresenters(srv *handler.Server) {
}
