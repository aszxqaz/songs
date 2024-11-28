package repository

import (
	"context"
	"songs/internal/common/pagination"
	"songs/internal/songs/domain"
	"time"
)

type FindSearchParams struct {
	Group       string
	Song        string
	Text        string
	ReleaseDate string
	Link        string
	Before      time.Time
	After       time.Time
}

type SearchSongsResult struct {
	Songs []*domain.Song
	Total int
}

type SongRepository interface {
	GetByID(ctx context.Context, ID int) (*domain.Song, error)
	Find(ctx context.Context, searchParams FindSearchParams, pagParams pagination.Params) (*SearchSongsResult, error)
	UpdateSong(ctx context.Context, ID int, updateFn func(s *domain.Song) (*domain.Song, error)) (*domain.Song, error)
	CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	DeleteByID(ctx context.Context, ID int) error
}
