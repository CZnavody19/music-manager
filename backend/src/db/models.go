package db

import "github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/model"

type TrackWithISRCs struct {
	model.Tracks
	ISRCs []model.TrackIsrcs
}
