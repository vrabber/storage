package repository

import (
	"context"

	"github.com/vrabber/storage/internal/models"
)

func (i *Implementation) InsertFileUpload(ctx context.Context, fileUpload *models.FileUpload) error {
	const query = `
		INSERT INTO file_uploads(upload_id, file_info_id, name, path) 
		VALUES ($1, $2, $3, $4)`

	_, err := i.db.Exec(
		ctx,
		query,
		fileUpload.UploadID,
		fileUpload.FileInfo.ID,
		fileUpload.Name,
		fileUpload.Path)
	return err
}

func (i *Implementation) GetFileUpload(ctx context.Context, uploadID string) (*models.FileUpload, error) {
	const query = `SELECT 
    				fu.id, 
    				fu.upload_id, 
    				fu.name, 
    				fu.path, 
    				fu.created_at, 
    				fu.updated_at,
    				fi.id,
    				fi.name,
    				fi.size,
    				fi.hash,
    				fi.owner,
    				fi.created_at,
    				fi.updated_at
					FROM file_uploads fu 
					JOIN file_info fi ON fi.id = fu.file_info_id
					WHERE fu.upload_id = $1`

	result := models.FileUpload{FileInfo: &models.FileInfo{}}
	err := i.db.QueryRow(ctx, query, uploadID).
		Scan(
			&result.ID,
			&result.UploadID,
			&result.Name,
			&result.Path,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.FileInfo.ID,
			&result.FileInfo.Name,
			&result.FileInfo.Size,
			&result.FileInfo.Hash,
			&result.FileInfo.Owner,
			&result.FileInfo.CreatedAt,
			&result.FileInfo.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
