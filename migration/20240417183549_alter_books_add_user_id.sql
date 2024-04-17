-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD user_id VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP COLUMN user_id;
-- +goose StatementEnd
