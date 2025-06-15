-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    feed_id UUID NOT NULL REFERENCES feed(id),
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;