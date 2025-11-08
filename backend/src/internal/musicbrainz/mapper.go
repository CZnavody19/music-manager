package musicbrainz

import (
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/google/uuid"
	"go.uploadedlobster.com/mbtypes"
	"go.uploadedlobster.com/musicbrainzws2"
)

func mapID(id mbtypes.MBID) (uuid.UUID, error) {
	strId := string(id)
	uid, err := uuid.Parse(strId)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}

func mapISRCs(isrcs []mbtypes.ISRC) []string {
	var out []string

	for _, isrc := range isrcs {
		out = append(out, isrc.Compact())
	}

	return out
}

func mapTrack(recording musicbrainzws2.Recording) (domain.Track, error) {
	id, err := mapID(recording.ID)
	if err != nil {
		return domain.Track{}, err
	}

	return domain.Track{
		ID:     id,
		Title:  recording.Title,
		Artist: recording.ArtistCredit.String(),
		Length: recording.Length.Milliseconds(),
		ISRCs:  mapISRCs(recording.ISRCs),
	}, nil
}
