-- name: GetVerses :many
SELECT
    verse_number,
    text
FROM
    verses
WHERE
    song_id = $1
ORDER BY
    verse_number
LIMIT $2 OFFSET $3;
