-- +goose Up
ALTER TABLE feed
ADD COLUMN created_at TIMESTAMP NOT NULL,
ADD COLUMN updated_at TIMESTAMP NOT NULL;

-- +goose Down
ALTER TABLE feed
DROP COLUMN created_at,
DROP COLUMN updated_at;