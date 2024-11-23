package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/asp3cto/songshelf/internal/data/repository/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
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

	db := sqlc.New(conn)
	data, err := db.GetVerses(ctx, sqlc.GetVersesParams{
		SongID: 1,
		Limit:  1,
		Offset: 0,
	})

	err = db.DeleteSong(ctx, 2)
	fmt.Println(err)

	fmt.Println(data)

	err = db.UpdateSong(ctx, sqlc.UpdateSongParams{
		ID:       1,
		Title:    "Supermassived Black Hole",
		ArtistID: 1,
		ReleaseDate: pgtype.Date{
			Time:  time.Date(2006, 7, 16, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	})

	fmt.Println(err)
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
