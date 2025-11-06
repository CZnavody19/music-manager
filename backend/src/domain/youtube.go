package domain

type YouTubeConfig struct {
	Enabled    bool
	OAuth      []byte
	Token      []byte
	PlaylistID string
}

type YouTubeVideo struct {
	VideoID       string
	Title         string
	ChannelTitle  string
	ThumbnailURL  *string
	Duration      *int64
	Position      int64
	NextPageToken string
}
