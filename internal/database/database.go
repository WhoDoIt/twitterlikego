package database

import (
	"context"
	"os"

	"github.com/WhoDoIt/twitterlikego/internal/logging"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Query(sql string, args ...any) (pgx.Rows, error) {
	if args != nil {
		return conn.Query(context.Background(), sql, args)
	} else {
		return conn.Query(context.Background(), sql)
	}
}

func QueryRow(sql string, args ...any) pgx.Row {
	if args != nil {
		return conn.QueryRow(context.Background(), sql, args)
	} else {
		return conn.QueryRow(context.Background(), sql)
	}
}

func init() {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_TWITTERLIKEGO"))
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}
	logging.Log("Connected to database")
}
