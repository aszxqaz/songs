package gorm

import (
	"context"
	"log/slog"
	"songs/internal/common/custom_error"
	"songs/internal/songs/domain"
	"strings"
)

// CreateSong implements song.Repository.
func (r *PgGormSongRepository) CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	songModel := newSongModelFromDomain(song)
	slog.Info("create song", "date", songModel.ReleaseDate.String())
	result := r.db.Create(songModel)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint \"idx_artist_title\"") {
			return nil, custom_error.NewConflictError(result.Error, "song with these artist and title already exists")
		}
		return nil, result.Error
	}

	return songModel.toDomain(), nil
}
