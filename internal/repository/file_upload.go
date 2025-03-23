package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/vrabber/storage/internal/models"
)

func (i *Implementation) InsertFileUpload(ctx context.Context, fileUpload *models.FileUpload) error {
	const query = `
		INSERT INTO file_uploads(upload_id, file_info_id) 
		VALUES ($1, $2)`

	return i.tx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		_, err := tx.Exec(
			ctx,
			query,
			fileUpload.UploadID,
			fileUpload.FileInfo.ID)
		return err
	})
}
