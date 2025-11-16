package youtube

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/CZnavody19/music-manager/src/db/youtube"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/internal/tidal"
	"github.com/CZnavody19/music-manager/src/utils"
	"github.com/google/uuid"
	"go.uploadedlobster.com/musicbrainzws2"
)

const (
	maxTimeDiff    = 5 // seconds
	titleWeight    = 0.5
	artistWeight   = 0.2
	durationWeight = 0.2
	isrcWeight     = 0.1
)

type IdentificationRequest struct {
	Video   *domain.YouTubeVideo
	YtStore *youtube.YouTubeStore
	Tidal   *tidal.Tidal
}

func (r IdentificationRequest) GetSearchQuery() string {
	at := extractArtistTitle(r.Video.Title, r.Video.ChannelTitle)

	qDur := *r.Video.Duration / 2

	var featuringQ string
	for _, artist := range at.Featuring {
		featuringQ += fmt.Sprintf("\"%s\" ", artist)
	}

	return fmt.Sprintf("artist:(+\"%s\" %s) AND recording:\"%s\" AND qdur:[%d TO %d]", at.Artist, featuringQ, at.Title, qDur-maxTimeDiff, qDur+maxTimeDiff)
}

func (r IdentificationRequest) GetSimilarityScore(recording *musicbrainzws2.Recording) float64 {
	at := extractArtistTitle(r.Video.Title, r.Video.ChannelTitle)

	tSim := utils.StringSimilarity(at.Title, recording.Title)
	aSim := utils.StringSimilarity(fmt.Sprintf("%s %s", at.Artist, strings.Join(at.Featuring, " ")), recording.ArtistCredit.String())
	dSim := utils.NumberSimilarity(float64(*r.Video.Duration), recording.Length.Seconds())
	iSim := 0.0
	if len(recording.ISRCs) > 0 {
		iSim = 1.0
	}

	return (tSim * titleWeight) + (aSim * artistWeight) + (dSim * durationWeight) + (iSim * isrcWeight)
}

func (r IdentificationRequest) LinkTrack(ctx context.Context, id uuid.UUID) error {
	return r.YtStore.LinkTrack(ctx, r.Video.VideoID, id)
}

type parsed struct {
	Artist    string
	Title     string
	Featuring []string
}

func extractFeaturing(raw string) (clean string, featured []string) {
	re := regexp.MustCompile(`(?i)(?:\(?\s*(?:ft\.?|feat\.?|featuring)\s+([^\)\-–]+)\)?)`)

	matches := re.FindAllStringSubmatch(raw, -1)
	for _, m := range matches {
		if len(m) > 1 {
			parts := regexp.MustCompile(`\s*(?:,|&|and)\s*`).Split(m[1], -1)
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p != "" {
					featured = append(featured, p)
				}
			}
		}
	}

	clean = re.ReplaceAllString(raw, "")
	clean = strings.TrimSpace(clean)
	clean = utils.CollapseSpaces(clean)

	return
}

func cleanTitle(raw string) string {
	s := utils.RemoveParentheses(raw)
	s = utils.NormalizeHyphens(s)
	s = utils.CollapseSpaces(s)
	s = strings.TrimSpace(s)
	return s
}

func getArtistName(channelName string) string {
	cleaned := cleanTitle(channelName)

	if strings.HasSuffix(cleaned, " - Topic") {
		return strings.TrimSpace(strings.TrimSuffix(cleaned, " - Topic"))
	}

	return cleaned
}

func parseArtistTitle(cleaned string) *parsed {
	if m := regexp.MustCompile(`^(.+?)\s*-\s*(.+)$`).FindStringSubmatch(cleaned); m != nil {
		return &parsed{
			Artist: strings.TrimSpace(m[1]),
			Title:  strings.TrimSpace(m[2]),
		}
	}
	return nil
}

func extractArtistTitle(rawTitle string, channelName string) parsed {
	title, featured := extractFeaturing(rawTitle)

	cleaned := cleanTitle(title)

	if strings.HasSuffix(channelName, " - Topic") {
		return parsed{
			Artist:    getArtistName(channelName),
			Title:     cleaned,
			Featuring: featured,
		}
	}

	if parsed := parseArtistTitle(cleaned); parsed != nil {
		parsed.Featuring = featured
		return *parsed
	}

	return parsed{
		Artist:    getArtistName(channelName),
		Title:     cleaned,
		Featuring: featured,
	}
}
