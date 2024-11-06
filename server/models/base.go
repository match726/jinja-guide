package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/xid"
)

type Postgres struct {
	dbPool *pgxpool.Pool
}

var pgInstance *Postgres

func NewPool() (*Postgres, error) {

	dbname := os.Getenv("POSTGRES_DATABASE")
	dsn := os.Getenv("POSTGRES_URL")

	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	} else {
		pgInstance = &Postgres{pool}
	}

	if err = pgInstance.dbPool.Ping(ctx); err != nil {
		return nil, err
	} else {
		fmt.Printf("Message: データベース[%s]へ接続", dbname)
	}

	return pgInstance, err

}

func (pg *Postgres) ClosePool() {
	pg.dbPool.Close()
}

// XIDの取得
func GetXID() (uid string) {

	uid = xid.New().String()
	return uid

}
