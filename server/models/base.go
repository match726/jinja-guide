package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/xid"
)

type Postgres struct {
	dbPool *pgxpool.Pool
}

var pgInstance *Postgres

// コネクションプールの作成
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
		fmt.Printf("NewPool: データベース[%s]へ接続\n", dbname)
	}

	return pgInstance, err

}

// コネクションプールのクローズ
func (pg *Postgres) ClosePool() {
	pg.dbPool.Close()
}

func GetNowTime() time.Time {

	var current time.Time

	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)
	current = time.Now().In(jstZone)

	return current

}

// XIDの取得
func GetXID() (uid string) {

	uid = xid.New().String()
	return uid

}
