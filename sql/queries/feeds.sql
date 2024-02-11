-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: FollowByFeedId :one
INSERT INTO feeds_users (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedById :one
SELECT * FROM feeds
WHERE id = $1;

-- name: DeleteFeedFollowByID :exec
DELETE FROM feeds_users
WHERE id = $1 AND user_id = $2;

-- name: GetFeedFollowById :one
SELECT * FROM feeds_users
WHERE id = $1;