package config

import (
	"context"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
)

func (cs *ConfigStore) GetYoutubeConfig(ctx context.Context) (*domain.YouTubeConfig, error) {
	stmt := table.Youtube.SELECT(table.Youtube.AllColumns).WHERE(table.Youtube.Active.IS_TRUE())

	var dest model.Youtube
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapYouTubeConfig(&dest), nil
}

func (cs *ConfigStore) SaveYoutubeFiles(ctx context.Context, oauthData, tokenData []byte) error {
	stmt := table.Youtube.INSERT(table.Youtube.OAuth, table.Youtube.Token, table.Youtube.Active).VALUES(oauthData, tokenData, true)

	stmt = db.DoUpsert(stmt, table.Youtube.Active, table.Youtube.MutableColumns, table.Youtube.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) SetYoutubeEnabled(ctx context.Context, enabled bool) error {
	stmt := table.Youtube.UPDATE(table.Youtube.Enabled).SET(enabled).WHERE(table.Youtube.Active.IS_TRUE())

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
