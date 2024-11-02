package models

import (
	"context"
	"fmt"
	"os"

	"github.com/goark/gocli/exitcode"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/xid"
)

var dbPool *pgxpool.Pool
var err error

func Run() exitcode.ExitCode {

	dbname := os.Getenv("POSTGRES_DATABASE")
	dsn := os.Getenv("POSTGRES_URL")

	dbPool, err := GetPool(context.Background(), dsn)

	if err != nil {
		fmt.Println(err)
		return exitcode.Abnormal
	} else {
		fmt.Printf("Message: データベース[%s]へ接続", dbname)
	}

	defer dbPool.Close()

	return exitcode.Normal

}

func GetPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {

	conn, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, conn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil

}

// XIDの取得
func GetXID() (uid string) {

	uid = xid.New().String()
	return uid

}
