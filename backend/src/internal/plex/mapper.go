package plex

import (
	"strings"

	"github.com/google/uuid"
)

func mapMbid(input string) (*uuid.UUID, error) {
	if input == "" || !strings.HasPrefix(input, "mbid://") {
		return nil, nil
	}

	id, err := uuid.Parse(strings.TrimPrefix(input, "mbid://"))
	if err != nil {
		return nil, err
	}

	return &id, nil
}
