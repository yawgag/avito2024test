package db

import (
	"avito2024test/config"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDb() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	Pool, err := pgxpool.New(context.Background(), cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	Pool.Close()
}
