package db

//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://music:music@localhost:5432/musicdb?sslmode=disable -schema=public -path=./gen
//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://music:music@localhost:5432/musicdb?sslmode=disable -schema=config -path=./gen
