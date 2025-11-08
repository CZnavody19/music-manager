package domain

import "go.uploadedlobster.com/musicbrainzws2"

type IdentificationRequest interface {
	GetSearchQuery() string
	GetSimilarityScore(recording *musicbrainzws2.Recording) float64
}
