package requests

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CZnavody19/music-manager/plexapi/options"
)

type ReqClient struct {
	httpClient *http.Client
	options    *options.ClientOptions
}

func NewReqClient(opt *options.ClientOptions) *ReqClient {
	return &ReqClient{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		options: opt,
	}
}

func (req *ReqClient) GET(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%d%s", req.options.Protocol, req.options.Host, req.options.Port, path)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-Plex-Token", req.options.Token)

	res, err := req.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	return body, nil
}

func (req *ReqClient) POST(ctx context.Context, path string) error {
	url := fmt.Sprintf("%s://%s:%d%s", req.options.Protocol, req.options.Host, req.options.Port, path)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-Plex-Token", req.options.Token)
	request.Header.Set("Content-Type", "application/json")

	_, err = req.httpClient.Do(request)
	if err != nil {
		return err
	}

	return nil
}
