package domain

import "github.com/google/uuid"

type PlexConfig struct {
	Enabled   bool
	Protocol  string
	Host      string
	Port      int64
	Token     string
	LibraryID int64
}

type PlexTrack struct {
	ID       int64
	Title    string
	Artist   string
	Duration int64
	Mbid     *uuid.UUID
	TrackID  *string
}
