package musicbrainz

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/music-manager/src/db/musicbrainz"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
	"github.com/CZnavody19/music-manager/src/internal/websockets"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uploadedlobster.com/musicbrainzws2"
)

type MusicBrainz struct {
	client      *musicbrainzws2.Client
	mbStore     *musicbrainz.MusicbrainzStore
	websockets  *websockets.Websockets
	SearchQueue chan domain.IdentificationRequest
	MatchQueue  chan domain.MatchRequest
}

func NewMusicBrainz(mbs *musicbrainz.MusicbrainzStore, ws *websockets.Websockets) (*MusicBrainz, error) {
	client := musicbrainzws2.NewClient(musicbrainzws2.AppInfo{
		Name:    "music-manager",
		Version: "0.1",
		URL:     "github.com/CZnavody19/music-manager",
	})

	mb := &MusicBrainz{
		client:      client,
		mbStore:     mbs,
		websockets:  ws,
		SearchQueue: make(chan domain.IdentificationRequest, 100),
		MatchQueue:  make(chan domain.MatchRequest, 100),
	}

	ctx := context.Background()

	go mb.searchWorker(ctx)
	go mb.matchWorker(ctx)

	return mb, nil
}

func (mb *MusicBrainz) GetTracks(ctx context.Context, notDownloaded bool) ([]*domain.Track, error) {
	return mb.mbStore.GetTracks(ctx, notDownloaded)
}

func (mb *MusicBrainz) DeleteTrack(ctx context.Context, id uuid.UUID) error {
	return mb.mbStore.DeleteTrack(ctx, id)
}

func (mb *MusicBrainz) searchWorker(ctx context.Context) {
	zap.S().Info("MusicBrainz search worker started")

	for request := range mb.SearchQueue {
		start := time.Now()

		mb.websockets.SendTask(&model.Task{
			Title:     "Processing MusicBrainz search request",
			StartedAt: start,
			Ended:     false,
		})

		zap.S().Info("Processing MusicBrainz search request")

		searchStr := request.GetSearchQuery()
		zap.S().Infof("MusicBrainz search query: %s", searchStr)

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

		track, err := mapTrack(most)
		if err != nil {
			zap.S().Errorf("Error mapping MusicBrainz track %s: %v", most.ID, err)
			continue
		}

		err = mb.mbStore.StoreTrack(ctx, track)
		if err != nil {
			zap.S().Errorf("Error storing MusicBrainz track %s: %v", most.ID, err)
			continue
		}

		err = request.LinkTrack(ctx, track.ID)
		if err != nil {
			zap.S().Errorf("Error linking MusicBrainz track %s: %v", most.ID, err)
			continue
		}

		mb.websockets.SendTask(&model.Task{
			Title:     "Processing MusicBrainz search request",
			StartedAt: start,
			Ended:     true,
		})

		zap.S().Infof("Stored MusicBrainz track: %s (similarity: %.2f)", most.ID, mostSim)
	}
}

func (mb *MusicBrainz) matchWorker(ctx context.Context) {
	zap.S().Info("MusicBrainz match worker started")

	for request := range mb.MatchQueue {
		start := time.Now()

		mb.websockets.SendTask(&model.Task{
			Title:     "Processing MusicBrainz match request",
			StartedAt: start,
			Ended:     false,
		})

		zap.S().Info("Processing MusicBrainz match request")

		filter := musicbrainzws2.SearchFilter{
			Query:    fmt.Sprintf("tid:\"%s\"", request.GetTrackID()),
			Includes: []string{"isrcs"},
		}
		paginator := musicbrainzws2.Paginator{
			Offset: 0,
			Limit:  100,
		}

		res, err := mb.client.SearchRecordings(ctx, filter, paginator)
		if err != nil {
			zap.S().Errorf("MusicBrainz match search error: %v", err)
			continue
		}

		track, err := mapTrack(res.Recordings[0])
		if err != nil {
			zap.S().Errorf("Error mapping MusicBrainz track %s: %v", res.Recordings[0].ID, err)
			continue
		}

		err = mb.mbStore.StoreTrack(ctx, track)
		if err != nil {
			zap.S().Errorf("Error storing MusicBrainz track %s: %v", res.Recordings[0].ID, err)
			continue
		}

		err = request.LinkTrack(ctx, track.ID)
		if err != nil {
			zap.S().Errorf("Error linking MusicBrainz track %s: %v", res.Recordings[0].ID, err)
			continue
		}

		mb.websockets.SendTask(&model.Task{
			Title:     "Processing MusicBrainz match request",
			StartedAt: start,
			Ended:     true,
		})

		zap.S().Infof("Stored MusicBrainz track: %s", res.Recordings[0].ID)
	}
}
