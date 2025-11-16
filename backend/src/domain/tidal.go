package domain

import "time"

type TidalConfig struct {
	Enabled          bool
	AuthTokenType    string
	AuthAccessToken  string
	AuthRefreshToken string
	AuthExpiresAt    time.Time
	AuthClientID     string
	AuthClientSecret string
	DownloadTimeout  int64
	DownloadRetries  int64
	DownloadThreads  int64
	AudioQuality     string
}
