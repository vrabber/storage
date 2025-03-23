package repository

import (
	"context"

	"github.com/vrabber/storage/internal/models"
)

func (i *Implementation) InsertFileInfo(ctx context.Context, fileInfo *models.FileInfo) error {
	const query = `
		INSERT INTO file_info(name, size, hash, owner, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	row := i.db.QueryRow(
		ctx,
		query,
		fileInfo.Name,
		fileInfo.Size,
		fileInfo.Hash,
		fileInfo.Owner,
		fileInfo.CreatedAt,
		fileInfo.UpdatedAt)
	return row.Scan(&fileInfo.ID)
}
