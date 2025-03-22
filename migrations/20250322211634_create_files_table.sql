-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files
(
    id           SERIAL       NOT NULL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    path         VARCHAR(255) NOT NULL,
    file_info_id INT          NOT NULL UNIQUE REFERENCES file_info (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP CHECK ( updated_at >= created_at ),
    UNIQUE (name, path)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
