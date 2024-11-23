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
    -- Attempt to insert the artist, or select the existing artist if already present
    INSERT INTO artists (name)
        VALUES ($1)
        ON CONFLICT (name) DO NOTHING
        RETURNING id
)
-- Select the artist ID, either from the insert or existing
INSERT INTO songs (title, artist_id)
SELECT $2, COALESCE(
        (SELECT id FROM inserted_artist),  -- Try to use the inserted artist's ID
        (SELECT id FROM artists WHERE artists.name = $1)  -- If no new artist was inserted, use the existing one
           )
RETURNING id;
