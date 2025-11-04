package http

import (
	"io"
	"net/http"

	"github.com/CZnavody19/music-manager/src/db/config"
	"go.uber.org/zap"
)

type HttpHandler struct {
	configStore *config.ConfigStore
}

func NewHttpHandler(cs *config.ConfigStore) *HttpHandler {
	return &HttpHandler{
		configStore: cs,
	}
}

func (hh *HttpHandler) UploadYoutubeCreds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	zap.L().Info("Processing upload youtube creds request")

	r.ParseMultipartForm(10 << 20) // 10 MB

	oauthFile, _, err := r.FormFile("oauth")
	if err != nil {
		zap.L().Error("Error retrieving the file", zap.Error(err))
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer oauthFile.Close()

	tokenFile, _, err := r.FormFile("token")
	if err != nil {
		zap.L().Error("Error retrieving the file", zap.Error(err))
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer tokenFile.Close()

	oauthBytes, err := io.ReadAll(oauthFile)
	if err != nil {
		zap.L().Error("Error reading the bytes", zap.Error(err))
		http.Error(w, "Error reading the bytes", http.StatusBadRequest)
		return
	}

	tokenBytes, err := io.ReadAll(tokenFile)
	if err != nil {
		zap.L().Error("Error reading the bytes", zap.Error(err))
		http.Error(w, "Error reading the bytes", http.StatusBadRequest)
		return
	}

	err = hh.configStore.SaveYoutubeFiles(ctx, oauthBytes, tokenBytes)
	if err != nil {
		zap.L().Error("Error saving youtube files", zap.Error(err))
		http.Error(w, "Error saving youtube files", http.StatusInternalServerError)
		return
	}

	zap.L().Info("Successfully uploaded youtube creds")
	w.WriteHeader(http.StatusOK)
}
