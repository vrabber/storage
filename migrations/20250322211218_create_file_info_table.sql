-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file_info
(
    id         SERIAL        NOT NULL PRIMARY KEY,
    name       VARCHAR(2048) NOT NULL,
    size       BIGINT        NOT NULL CHECK ( SIZE >= 0 ),
    hash       VARCHAR(255)  NOT NULL,
    owner      VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_info;
-- +goose StatementEnd
