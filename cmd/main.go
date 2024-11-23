package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"github.com/asp3cto/songshelf/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("pg config: %v", err)
	}

	conn, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("connection: %v", err)
	}
	defer conn.Close()

	if err := migrate(pgConfig.DSN()); err != nil {
		log.Fatal("migrate: ", err)
	}

}

func migrate(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	if err := goose.Up(db, "sql/migrations"); err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
