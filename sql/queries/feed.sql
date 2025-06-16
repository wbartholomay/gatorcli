-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feed;

-- name: GetFeedByUrl :one
SELECT * FROM feed
WHERE feed.url = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feed;

-- name: MarkFeedFetched :exec
UPDATE feed
SET last_fetched_at = $2, updated_at = $3
WHERE feed.id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feed
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;