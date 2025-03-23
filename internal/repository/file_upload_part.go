package repository

import (
	"context"
	"fmt"

	"github.com/vrabber/storage/internal/models"
)

func (i *Implementation) InsertFileUploadPart(ctx context.Context, part *models.FileUploadPart) error {
	const query = `INSERT INTO file_upload_parts(file_upload_id, size, offset_, started_at, finished_at) 
					SELECT id, $1, $2, $3, $4 FROM file_uploads WHERE upload_id = $5`
	_, err := i.db.Exec(ctx, query, part.Size, part.Offset, part.StartedAt, part.FinishedAt, part.FileUpload.UploadID)
	return err
}

func (i *Implementation) UpdateFileUploadPartFinished(ctx context.Context, part *models.FileUploadPart) error {
	const query = `UPDATE file_upload_parts SET finished_at = $1 WHERE file_upload_id = (SELECT id FROM file_uploads WHERE upload_id = $2)`
	r, err := i.db.Exec(ctx, query, part.FinishedAt, part.FileUpload.UploadID)
	if err == nil && r.RowsAffected() != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", r.RowsAffected())
	}
	return err
}

func (i *Implementation) IsFullLoaded(ctx context.Context, fileUpload *models.FileUpload) (bool, error) {
	//todo
	return true, nil
}
