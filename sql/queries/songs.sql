-- name: DeleteSong :exec
DELETE FROM songs WHERE id = $1;

-- name: UpdateSong :exec
UPDATE songs
SET
    artist_id = $2,
    title = $3,
    release_date = $4
WHERE
    id = $1;

-- name: InsertSong :one
WITH inserted_artist AS (
    INSERT INTO artists (name)
        VALUES ($1)
        ON CONFLICT (name) DO NOTHING
        RETURNING id
)
INSERT INTO songs (title, artist_id)
SELECT $2, COALESCE(
    (SELECT id FROM inserted_artist),
    (SELECT id FROM artists WHERE artists.name = $1)
)
RETURNING id;
