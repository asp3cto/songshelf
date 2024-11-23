-- name: DeleteSong :exec
DELETE FROM songs WHERE id = $1;