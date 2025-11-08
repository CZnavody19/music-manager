package areas

import (
	"context"
	"fmt"

	"github.com/CZnavody19/music-manager/plexapi/requests"
)

type Library struct {
	reqClient *requests.ReqClient
}

func NewLibrary(reqClient *requests.ReqClient) *Library {
	return &Library{
		reqClient: reqClient,
	}
}

func (l *Library) RefreshSection(ctx context.Context, sectionID int) error {
	return l.reqClient.POST(ctx, fmt.Sprintf("/library/sections/%d/refresh", sectionID))
}
