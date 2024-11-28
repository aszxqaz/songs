package gorm

import (
	"context"
	"errors"
	"songs/internal/songs/domain"

	"gorm.io/gorm"
)

// GetByID implements song.Repository.
func (r *PgGormSongRepository) GetByID(ctx context.Context, ID int) (*domain.Song, error) {
	var songModel songModel
	result := r.db.First(&songModel, ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return songModel.toDomain(), nil
}
