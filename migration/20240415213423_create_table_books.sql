-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books
(
    id                    VARCHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    author                VARCHAR(255) NOT NULL,
    isbn                  VARCHAR(50) NOT NULL,
    publisher             VARCHAR(255) NOT NULL,
    publication_date      TIMESTAMP NOT NULL,
    edition               VARCHAR(50) NULL,
    genre                 VARCHAR(50) NOT NULL,
    language              VARCHAR(50) NOT NULL,
    number_of_pages       INT NOT NULL,
    description           TEXT NOT NULL,
    price                 DECIMAL(10, 2) NULL,
    format                VARCHAR(50) NOT NULL,

    created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at            TIMESTAMP NULL DEFAULT NULL,
    is_not_archived       BOOLEAN GENERATED ALWAYS AS (IF(deleted_at IS NULL, 1, null)) VIRTUAL COMMENT 'flag for soft delete'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
