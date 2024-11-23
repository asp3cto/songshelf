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