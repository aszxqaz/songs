package gorm

import (
	"context"
	"songs/internal/common/date"
	"songs/internal/common/pagination"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/domain"
)

// Find implements song.Repository.
func (r *PgGormSongRepository) Find(
	ctx context.Context,
	searchParams repository.FindSearchParams,
	pagParams pagination.Params,
) (*repository.SearchSongsResult, error) {
	var gormSongs []*songModel
	var total int64

	result := r.db.
		Where("artist ILIKE ?", "%"+searchParams.Group+"%").
		Where("title ILIKE ?", "%"+searchParams.Song+"%").
		Where("lyrics ILIKE ?", "%"+searchParams.Text+"%").
		Where("link ILIKE ?", "%"+searchParams.Link+"%").
		Where("release_date >= ?", date.TimeToDate(searchParams.After)).
		Where("release_date <= ?", date.TimeToDate(searchParams.Before)).
		Find(&gormSongs).
		Offset(pagParams.Offset).
		Limit(pagParams.Limit).
		Count(&total).
		Order("id ASC")

	if result.Error != nil {
		return nil, result.Error
	}

	songs := make([]*domain.Song, len(gormSongs))
	for i, gormSong := range gormSongs {
		songs[i] = gormSong.toDomain()
	}

	return &repository.SearchSongsResult{
		Songs: songs,
		Total: int(total),
	}, nil
}
