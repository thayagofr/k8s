package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func OpenConnection(ctx context.Context, DSN string) (*sql.DB, error) {
	db, err = sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}
	return db, db.PingContext(ctx)
}
