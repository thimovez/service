-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments (
    id text PRIMARY KEY NOT NULL,
    user_id text NOT NULL,
    content text NOT NULL,
    parent_id text NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES comments(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd