package musicbrainz

import (
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/model"
	"github.com/CZnavody19/music-manager/src/domain"
)

func mapISRCs(track domain.Track) []model.TrackIsrcs {
	var out []model.TrackIsrcs

	for _, isrc := range track.ISRCs {
		out = append(out, model.TrackIsrcs{
			TrackID: track.ID,
			Isrc:    isrc,
		})
	}

	return out
}
