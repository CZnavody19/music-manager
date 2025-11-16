package setup

import (
	"github.com/CZnavody19/music-manager/src/graph"
	"github.com/gorilla/mux"
)

func SetupHTTPHandlers(router *mux.Router, resolver *graph.Resolver) {
	router.HandleFunc("/upload/youtube", resolver.HttpHandler.UploadYoutubeCreds).Methods("POST")

	router.HandleFunc("/config/tidal", resolver.HttpHandler.GetTidalConfig).Methods("GET")
}
