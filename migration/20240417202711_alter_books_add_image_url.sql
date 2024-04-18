-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD image_url TEXT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN image_url;
-- +goose StatementEnd
