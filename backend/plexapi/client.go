package plexapi

import (
	"github.com/CZnavody19/music-manager/plexapi/areas"
	"github.com/CZnavody19/music-manager/plexapi/options"
	"github.com/CZnavody19/music-manager/plexapi/requests"
)

type Client struct {
	Content *areas.Content
	Library *areas.Library

	reqClient *requests.ReqClient
	options   *options.ClientOptions
}

func NewClient(options options.ClientOptions) *Client {
	reqClient := requests.NewReqClient(&options)

	client := Client{
		reqClient: reqClient,
		options:   &options,
	}

	client.Content = areas.NewContent(reqClient)
	client.Library = areas.NewLibrary(reqClient)

	return &client
}
