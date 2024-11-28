package gorm

import (
	"context"
	"errors"
	custom_error "songs/internal/common/custom_error"
)

// GetByID implements song.Repository.
func (r *PgGormSongRepository) DeleteByID(ctx context.Context, ID int) error {
	result := r.db.Delete(&songModel{}, ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return custom_error.NewNotFoundError(
			errors.New("song not found"),
		)
	}

	return nil
}
