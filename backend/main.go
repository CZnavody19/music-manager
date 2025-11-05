package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/CZnavody19/music-manager/src/config"
	"github.com/CZnavody19/music-manager/src/graph"
	"github.com/CZnavody19/music-manager/src/graph/generated"
	"github.com/CZnavody19/music-manager/src/setup"
	"github.com/gorilla/mux"
	"github.com/nextap-solutions/goNextService"
	"github.com/nextap-solutions/goNextService/components"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
)

type ServerComponents struct {
	httpServer goNextService.Component
}

func main() {
	err := serve()
	if err != nil {
		fmt.Println("Error running the application: ", err)
		os.Exit(1)
	}
}

func serve() error {
	configuration := config.LoadConfig()

	api, err := setupService(configuration)
	if err != nil {
		return err
	}

	app := goNextService.NewApplications(api.httpServer)
	app.WithLogger(zap.S())

	return app.Run()
}

func setupService(configuration *config.Config) (*ServerComponents, error) {
	setup.InitLogger(*configuration)
	s, _ := json.MarshalIndent(configuration, "", "\t")
	zap.S().Info("Logger initialized successfully")
	zap.S().Info(string(s))

	dbConn, err := setup.SetupDb(&configuration.DBConfig)
	if err != nil {
		zap.S().Error("Error setting up db connection")
		zap.S().Error(err.Error())
		return nil, err
	}

	resolver, err := setup.NewResolver(dbConn, *configuration)
	if err != nil {
		zap.S().Error("Error setting up resolver")
		zap.S().Error(err.Error())
		return nil, err
	}

	c := generated.Config{
		Resolvers: resolver,
	}

	directives := graph.NewDirectives(resolver)

	setup.SetupDirectives(&c, directives)

	srv := handler.New(generated.NewExecutableSchema(c))
	srv.Use(graph.LoggingExtension{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	if configuration.Server.IntrospectionEnable {
		srv.Use(extension.Introspection{})
	}
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	setup.SetupPresenters(srv)

	router := mux.NewRouter()
	router.Use(resolver.Auth.Middleware())
	router.Handle("/", srv)
	if configuration.Server.PlaygroundEnable {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	}

	setup.SetupHTTPHandlers(router, resolver)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(router)
	api := http.Server{
		Addr:         "0.0.0.0:" + configuration.Server.Port,
		ReadTimeout:  configuration.Server.ReadTimeout,
		WriteTimeout: configuration.Server.WriteTimeout,
		Handler:      handler,
	}
	httpComponent := components.NewHttpComponent(handler, components.WithHttpServer(&api))

	return &ServerComponents{
		httpServer: httpComponent,
	}, nil
}
