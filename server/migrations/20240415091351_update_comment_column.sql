-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE comments
    DROP CONSTRAINT comments_parent_id_fkey;
-- +goose StatementEnd
