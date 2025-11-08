package musicbrainz

import (
	"context"

	"github.com/CZnavody19/music-manager/src/domain"
	"go.uber.org/zap"
	"go.uploadedlobster.com/musicbrainzws2"
)

type MusicBrainz struct {
	client      *musicbrainzws2.Client
	SearchQueue chan domain.IdentificationRequest
}

func NewMusicBrainz() (*MusicBrainz, error) {
	client := musicbrainzws2.NewClient(musicbrainzws2.AppInfo{
		Name:    "music-manager",
		Version: "0.1",
	})

	mb := &MusicBrainz{
		client:      client,
		SearchQueue: make(chan domain.IdentificationRequest, 100),
	}

	ctx := context.Background()

	go mb.searchWorker(ctx)

	return mb, nil
}

func (mb *MusicBrainz) searchWorker(ctx context.Context) {
	zap.S().Info("MusicBrainz search worker started")

	for request := range mb.SearchQueue {
		zap.S().Info("Processing MusicBrainz search request")

		searchStr := request.GetSearchQuery()

		filter := musicbrainzws2.SearchFilter{
			Query:    searchStr,
			Includes: []string{"isrcs"},
		}
		paginator := musicbrainzws2.Paginator{
			Offset: 0,
			Limit:  100,
		}

		res, err := mb.client.SearchRecordings(ctx, filter, paginator)
		if err != nil {
			zap.S().Errorf("MusicBrainz search error: %v", err)
			continue
		}

		var most musicbrainzws2.Recording
		mostSim := 0.0

		for _, recording := range res.Recordings {
			sim := request.GetSimilarityScore(&recording)

			if sim > mostSim {
				mostSim = sim
				most = recording
			}
		}

		zap.S().Infof("Most similar recording: %s (score: %f) [ID: %s]", most.Title, mostSim, most.ID)
	}
}
