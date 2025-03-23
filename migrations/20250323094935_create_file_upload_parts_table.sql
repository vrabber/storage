-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file_upload_parts
(
    id             SERIAL NOT NULL PRIMARY KEY,
    file_upload_id INT    NOT NULL REFERENCES file_uploads (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    size BIGINT NOT NULL CHECK ( size > 0 ),
    offset_ BIGINT NOT NULL CHECK ( offset_ >= 0 ),
    started_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP DEFAULT NULL CHECK ( finished_at IS NULL OR finished_at >= started_at )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_upload_parts;
-- +goose StatementEnd
