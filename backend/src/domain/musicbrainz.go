package domain

import (
	"github.com/google/uuid"
	"go.uploadedlobster.com/musicbrainzws2"
)

type IdentificationRequest interface {
	GetSearchQuery() string
	GetSimilarityScore(recording *musicbrainzws2.Recording) float64
}

type Track struct {
	ID     uuid.UUID
	Title  string
	Artist string
	Length int64
	ISRCs  []string
}
