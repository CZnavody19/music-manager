package http

import (
	"encoding/json"
	"net/http"

	"github.com/CZnavody19/music-manager/src/domain"
)

type tidalConfig struct {
	General *domain.GeneralConfig
	Tidal   *domain.TidalConfig
}

func (hh *HttpHandler) GetTidalConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	gConfig, err := hh.configStore.GetGeneralConfig(ctx)
	if err != nil {
		http.Error(w, "Failed to get General config: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tConfig, err := hh.configStore.GetTidalConfig(ctx)
	if err != nil {
		http.Error(w, "Failed to get Tidal config: "+err.Error(), http.StatusInternalServerError)
		return
	}

	config := &tidalConfig{
		General: gConfig,
		Tidal:   tConfig,
	}

	err = json.NewEncoder(w).Encode(config)
	if err != nil {
		http.Error(w, "Failed to encode Tidal config: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
