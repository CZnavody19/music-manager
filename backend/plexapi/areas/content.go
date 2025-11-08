package areas

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CZnavody19/music-manager/plexapi/requests"
	"github.com/CZnavody19/music-manager/plexapi/types"
)

type Content struct {
	reqClient *requests.ReqClient
}

func NewContent(reqClient *requests.ReqClient) *Content {
	return &Content{
		reqClient: reqClient,
	}
}

func (c *Content) GetSectionLeaves(ctx context.Context, sectionID int) (*types.ContentBodyStrGuid, error) {
	body, err := c.reqClient.GET(ctx, fmt.Sprintf("/library/sections/%d/allLeaves", sectionID))
	if err != nil {
		return nil, err
	}

	var contentBody types.ContentBodyStrGuid
	err = json.Unmarshal(body, &contentBody)
	if err != nil {
		return nil, err
	}

	return &contentBody, nil
}

func (c *Content) GetMetadataItem(ctx context.Context, ids []string) (*types.ContentBody, error) {
	body, err := c.reqClient.GET(ctx, fmt.Sprintf("/library/metadata/%s", strings.Join(ids, ",")))
	if err != nil {
		return nil, err
	}

	var contentBody types.ContentBody
	err = json.Unmarshal(body, &contentBody)
	if err != nil {
		return nil, err
	}

	return &contentBody, nil
}
