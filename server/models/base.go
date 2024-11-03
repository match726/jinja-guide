package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/xid"
)

var dbPool *pgxpool.Pool

func NewPool() (*pgxpool.Pool, error) {

	dbname := os.Getenv("POSTGRES_DATABASE")
	dsn := os.Getenv("POSTGRES_URL")

	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbPool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err = dbPool.Ping(ctx); err != nil {
		return nil, err
	} else {
		fmt.Printf("Message: データベース[%s]へ接続", dbname)
	}

	return dbPool, err

}

// XIDの取得
func GetXID() (uid string) {

	uid = xid.New().String()
	return uid

}
