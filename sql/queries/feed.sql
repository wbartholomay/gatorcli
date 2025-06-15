-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetFeeds :many
SELECT * FROM feed;

-- name: GetFeedByUrl :one
SELECT * FROM feed
WHERE feed.url = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feed;