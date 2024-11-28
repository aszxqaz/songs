package command

import (
	"context"
	"songs/internal/common/decorator"
	"songs/internal/songs/deps/repository"

	"github.com/sirupsen/logrus"
)

type DeleteSongCommand struct {
	ID int
}

type DeleteSongResult struct{}

type DeleteSongCommandHandler decorator.CommandHandler[DeleteSongCommand, *DeleteSongResult]

type deleteSongCommandHandler struct {
	songRepo repository.SongRepository
}

func NewDeleteSongCommandHandler(
	songRepo repository.SongRepository,
	logger *logrus.Entry,
) DeleteSongCommandHandler {
	return decorator.ApplyCommandDecorators(&deleteSongCommandHandler{songRepo}, logger)
}

// Handle implements DeleteSongCommandHandler.
func (h *deleteSongCommandHandler) Handle(ctx context.Context, cmd DeleteSongCommand) (*DeleteSongResult, error) {
	err := h.songRepo.DeleteByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	return &DeleteSongResult{}, nil
}
