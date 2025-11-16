package domain

import (
	"context"

	"github.com/google/uuid"
	"go.uploadedlobster.com/musicbrainzws2"
)

type IdentificationRequest interface {
	GetSearchQuery() string
	GetSimilarityScore(recording *musicbrainzws2.Recording) float64
	LinkTrack(ctx context.Context, trackID uuid.UUID) error
}

type MatchRequest interface {
	GetTrackID() uuid.UUID
	LinkTrack(ctx context.Context, trackID uuid.UUID) error
}

type Track struct {
	ID            uuid.UUID
	Title         string
	Artist        string
	Length        int64
	ISRCs         []string
	LinkedYoutube bool
	LinkedPlex    bool
}
