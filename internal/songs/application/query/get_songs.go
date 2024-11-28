package query

import (
	"context"
	"songs/internal/common/custom_error"
	"songs/internal/common/decorator"
	"songs/internal/common/pagination"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/domain"

	"github.com/sirupsen/logrus"
)

var getSongsPaginationConstraints = pagination.Constraints{
	MinLimit:  1,
	MaxLimit:  100,
	MaxOffset: 200,
}

var getSongsPaginationDefaults = pagination.Defaults{
	Limit:  2,
	Offset: 0,
}

type GetSongsQuery struct {
	SearchParams repository.FindSearchParams
	PagParams    pagination.ParamsOptional
}

type Pagination struct {
	Limit  int
	Offset int
	Total  int
}

type GetSongsQueryResult struct {
	Songs      []*domain.Song
	Pagination Pagination
}

type GetSongsQueryHandler decorator.QueryHandler[GetSongsQuery, *GetSongsQueryResult]

type getSongsQueryHandler struct {
	songRepo repository.SongRepository
}

func NewGetSongsQueryHandler(
	songRepo repository.SongRepository,
	logger *logrus.Entry,
) GetSongsQueryHandler {
	return decorator.ApplyQueryDecorators(
		getSongsQueryHandler{
			songRepo,
		},
		logger,
	)
}

// Handle implements GetSongsQueryHandler.
func (g getSongsQueryHandler) Handle(ctx context.Context, query GetSongsQuery) (*GetSongsQueryResult, error) {
	pagParams := query.PagParams.MergeDefaults(getSongsPaginationDefaults)
	err := pagParams.CheckConstraints(getSongsPaginationConstraints)
	if err != nil {
		return nil, err
	}

	r, err := g.songRepo.Find(ctx, query.SearchParams, pagParams)
	if err != nil {
		return nil, custom_error.NewInternalError(err, "Database error")
	}

	return &GetSongsQueryResult{
		Songs: r.Songs,
		Pagination: Pagination{
			Limit:  pagParams.Limit,
			Offset: pagParams.Offset,
			Total:  r.Total,
		},
	}, nil
}
