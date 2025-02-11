package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
)

type Postgres struct {
	DbPool *pgxpool.Pool
}

var pgInstance *Postgres
var dbname string = os.Getenv("POSTGRES_DATABASE")
var dsn string = os.Getenv("POSTGRES_URL")

// コネクションプールの作成
func NewPool(ctx context.Context) (*Postgres, error) {

	var cfg *pgxpool.Config
	var pool *pgxpool.Pool
	var err error

	cfg, err = pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("接続情報構成失敗: %w", err)
	}

	pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("コネクションプール作成失敗: %w", err)
	}

	pgInstance = &Postgres{pool}

	if err = pgInstance.DbPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("プール接続失敗(Ping): %w", err)
	}

	logger.Info(ctx, "コネクションプール作成成功", "dbname", dbname)
	return pgInstance, nil

}

// コネクションプールの切断
func (pg *Postgres) ClosePool(ctx context.Context) {

	pg.DbPool.Close()
	logger.Info(ctx, "コネクションプール切断成功", "dbname", dbname)

}

// 現在日時の取得
func GetNowTime() (current time.Time) {

	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)
	current = time.Now().In(jstZone)

	return current

}
