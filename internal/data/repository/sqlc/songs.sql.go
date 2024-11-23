// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: songs.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteSong = `-- name: DeleteSong :exec
DELETE FROM songs WHERE id = $1
`

func (q *Queries) DeleteSong(ctx context.Context, id int) error {
	_, err := q.db.Exec(ctx, deleteSong, id)
	return err
}

const insertSong = `-- name: InsertSong :one
WITH inserted_artist AS (
    -- Attempt to insert the artist, or select the existing artist if already present
    INSERT INTO artists (name)
        VALUES ($1)
        ON CONFLICT (name) DO NOTHING
        RETURNING id
)
INSERT INTO songs (title, artist_id)
SELECT $2, COALESCE(
        (SELECT id FROM inserted_artist),  -- Try to use the inserted artist's ID
        (SELECT id FROM artists WHERE artists.name = $1)  -- If no new artist was inserted, use the existing one
           )
RETURNING id
`

type InsertSongParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

// Select the artist ID, either from the insert or existing
func (q *Queries) InsertSong(ctx context.Context, arg InsertSongParams) (int, error) {
	row := q.db.QueryRow(ctx, insertSong, arg.Name, arg.Title)
	var id int
	err := row.Scan(&id)
	return id, err
}

const updateSong = `-- name: UpdateSong :exec
UPDATE songs
SET
    artist_id = $2,
    title = $3,
    release_date = $4
WHERE
    id = $1
`

type UpdateSongParams struct {
	ID          int         `json:"id"`
	ArtistID    int         `json:"artist_id"`
	Title       string      `json:"title"`
	ReleaseDate pgtype.Date `json:"release_date"`
}

func (q *Queries) UpdateSong(ctx context.Context, arg UpdateSongParams) error {
	_, err := q.db.Exec(ctx, updateSong,
		arg.ID,
		arg.ArtistID,
		arg.Title,
		arg.ReleaseDate,
	)
	return err
}
