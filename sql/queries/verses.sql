-- name: GetVersesBySongName :many
SELECT
    v.verse_number,
    v.text
FROM
    verses v
        JOIN
    songs s ON v.song_id = s.id
WHERE
    s.title = $1
ORDER BY
    v.verse_number
LIMIT $2 OFFSET $3;
