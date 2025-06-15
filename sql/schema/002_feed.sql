-- +goose Up
CREATE TABLE feed(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id)
        ON DELETE CASCADE,
    UNIQUE(url)
);

-- +goose Down
DROP TABLE feed;