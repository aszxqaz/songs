package gorm

import (
	"context"
	"errors"
	"songs/internal/common/custom_error"
	"songs/internal/songs/domain"

	"gorm.io/gorm"
)

// UpdateSong implements song.Repository.
func (r *PgGormSongRepository) UpdateSong(
	ctx context.Context,
	ID int,
	updateFn func(s *domain.Song) (*domain.Song, error),
) (*domain.Song, error) {
	var songModel songModel
	result := r.db.First(&songModel, ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, custom_error.NewNotFoundError(errors.New("song not found"), "Song not found")
		}
		return nil, result.Error
	}

	updatedSong, err := updateFn(songModel.toDomain())
	if err != nil {
		return nil, err
	}

	songModel.mergeWithDomain(updatedSong)
	result = r.db.Save(songModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return songModel.toDomain(), nil
}
