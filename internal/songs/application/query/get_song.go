package query

import (
	"context"
	"errors"
	custom_error "songs/internal/common/custom_error"
	"songs/internal/common/decorator"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/domain"

	"github.com/sirupsen/logrus"
)

type GetSongQuery struct {
	ID int
}

type GetSongQueryResult struct {
	Song *domain.Song
}

type GetSongQueryHandler decorator.QueryHandler[GetSongQuery, *GetSongQueryResult]

type getSongQueryHandler struct {
	songRepo repository.SongRepository
}

func NewGetSongQueryHandler(
	songRepo repository.SongRepository,
	logger *logrus.Entry,
) GetSongQueryHandler {
	return decorator.ApplyQueryDecorators(
		getSongQueryHandler{
			songRepo,
		},
		logger,
	)
}

// Handle implements GetSongQueryHandler.
func (g getSongQueryHandler) Handle(ctx context.Context, query GetSongQuery) (*GetSongQueryResult, error) {
	song, err := g.songRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	if song == nil {
		return nil, custom_error.NewNotFoundError(
			errors.New("song not found"),
		)
	}

	return &GetSongQueryResult{
		Song: song,
	}, nil
}
