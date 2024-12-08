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
var dbname string = os.Getenv("POSTGRES_DATABASE")

// コネクションプールの作成
func NewPool() (*Postgres, error) {

	var cfg *pgxpool.Config
	var pool *pgxpool.Pool
	var err error

	dsn := os.Getenv("POSTGRES_URL")
	ctx := context.Background()

	cfg, err = pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig(): %w", err)
	}

	pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig(): %w", err)
	}

	pgInstance = &Postgres{pool}

	if err = pgInstance.dbPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("dbPool.Ping(): %w", err)
	}

	fmt.Printf("NewPool: データベース[%s] 接続成功\n", dbname)
	return pgInstance, nil

}

// コネクションプールのクローズ
func (pg *Postgres) ClosePool() {

	pg.dbPool.Close()
	fmt.Printf("ClosePool: データベース[%s] 切断成功\n", dbname)

}

// 現在日時の取得
func GetNowTime() (current time.Time) {

	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)
	current = time.Now().In(jstZone)

	return current

}

// XIDの取得
func GetXID() (uid string) {

	uid = xid.New().String()
	return uid

}
