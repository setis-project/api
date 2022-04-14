package core

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

type Db struct {
	Ctx  *context.Context
	Conn *pgx.Conn
}

func DbConnect(ctx *context.Context) (*Db, error) {
	var (
		host     = os.Getenv("DB_HOST")
		database = os.Getenv("DB_NAME")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		port     = os.Getenv("DB_PORT")
		url      = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + database
	)

	db := Db{Ctx: ctx}
	conn, err := pgx.Connect(*db.Ctx, url)
	if err != nil {
		return &db, err
	}
	db.Conn = conn
	return &db, nil
}
