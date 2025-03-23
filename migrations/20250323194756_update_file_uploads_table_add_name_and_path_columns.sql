-- +goose Up
-- +goose StatementBegin
ALTER TABLE file_uploads ADD COLUMN name VARCHAR(255) NOT NULL;
ALTER TABLE file_uploads ADD COLUMN path VARCHAR(255) NOT NULL;
ALTER TABLE file_uploads ADD UNIQUE (name, path);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE file_uploads DROP CONSTRAINT file_uploads_name_path_key;
ALTER TABLE file_uploads DROP COLUMN IF EXISTS name;
ALTER TABLE file_uploads DROP COLUMN IF EXISTS path;
-- +goose StatementEnd
