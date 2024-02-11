-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE feeds_users (
    id UUID PRIMARY KEY,
    feed_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE(feed_id, user_id)
);

INSERT INTO feeds_users (id, feed_id, user_id, created_at, updated_at)
SELECT uuid_generate_v4(), f.id AS feed_id, u.id AS user_id, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
FROM feeds f
JOIN users u ON u.id = f.user_id;

-- +goose Down
DROP TABLE feeds_users;