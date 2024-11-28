package command

import (
	"context"
	"songs/internal/common/decorator"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/domain"

	"github.com/sirupsen/logrus"
)

type UpdateSongCommand struct {
	ID          int
	Artist      *string
	Title       *string
	Lyrics      *string
	ReleaseDate *string
	Link        *string
}

type UpdateSongResult struct {
	Song *domain.Song
}

type UpdateSongCommandHandler decorator.CommandHandler[UpdateSongCommand, *UpdateSongResult]

type updateSongCommandHandler struct {
	songRepo repository.SongRepository
}

func NewUpdateSongCommandHandler(songRepo repository.SongRepository, logger *logrus.Entry) UpdateSongCommandHandler {
	return decorator.ApplyCommandDecorators(&updateSongCommandHandler{songRepo}, logger)
}

// Handle implements UpdateSongCommandHandler.
func (h *updateSongCommandHandler) Handle(ctx context.Context, cmd UpdateSongCommand) (*UpdateSongResult, error) {
	song, err := h.songRepo.UpdateSong(ctx, cmd.ID, func(s *domain.Song) (*domain.Song, error) {
		err := s.Update(domain.UpdateSongData{
			Artist:      cmd.Artist,
			Title:       cmd.Title,
			Lyrics:      cmd.Lyrics,
			ReleaseDate: cmd.ReleaseDate,
			Link:        cmd.Link,
		})
		if err != nil {
			return nil, err
		}

		return s, nil
	})

	if err != nil {
		return nil, err
	}

	return &UpdateSongResult{Song: song}, nil
}
