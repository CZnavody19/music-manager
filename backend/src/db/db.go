package db

import (
	"github.com/go-jet/jet/v2/postgres"
)

//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://music:music@localhost:5432/musicdb?sslmode=disable -schema=public -path=./gen
//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://music:music@localhost:5432/musicdb?sslmode=disable -schema=config -path=./gen

type StringVal struct {
	Value string
}

func DoUpsert(stmt postgres.InsertStatement, conflictColumn postgres.Column, columns, excluded postgres.ColumnList) postgres.InsertStatement {
	return stmt.ON_CONFLICT(conflictColumn).DO_UPDATE(postgres.SET(columns.SET(excluded)))
}
