-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS images (
    id bigserial PRIMARY KEY NOT NULL,
    user_id bigserial NOT NULL,
    image_path text NOT NULL,
    image_url text NOT NULL
)
-- +goose StatementEnd

ALTER TABLE "images" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;
-- +goose StatementEnd