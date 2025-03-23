-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file_uploads
(
    id           SERIAL      NOT NULL PRIMARY KEY,
    upload_id    UUID UNIQUE NOT NULL,
    file_info_id INT         NOT NULL UNIQUE REFERENCES file_info (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP CHECK ( updated_at >= created_at )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_uploads;
-- +goose StatementEnd
