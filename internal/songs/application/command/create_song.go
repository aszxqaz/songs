package command

import (
	"context"
	"songs/internal/common/custom_error"
	"songs/internal/common/decorator"
	"songs/internal/songs/deps/info"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/domain"

	"github.com/sirupsen/logrus"
)

type CreateSongCommand struct {
	Artist string
	Title  string
}

type CreateSongResult struct {
	Song *domain.Song
}

type CreateSongCommandHandler decorator.CommandHandler[CreateSongCommand, *CreateSongResult]

type createSongCommandHandler struct {
	songRepo        repository.SongRepository
	songInfoService info.SongInfoAPIService
}

func NewCreateSongCommandHandler(
	songRepo repository.SongRepository,
	songInfoService info.SongInfoAPIService,
	logger *logrus.Entry,
) CreateSongCommandHandler {
	return decorator.ApplyCommandDecorators(&createSongCommandHandler{songRepo, songInfoService}, logger)
}

// Handle implements CreateSongCommandHandler.
func (h *createSongCommandHandler) Handle(ctx context.Context, cmd CreateSongCommand) (*CreateSongResult, error) {
	songInfo, err := h.songInfoService.GetSongInfo(ctx, info.Query{
		Artist: cmd.Artist,
		Title:  cmd.Title,
	})
	if err != nil {
		return nil, custom_error.NewInternalError(err, "failed to connect to song info service")
	}

	song := &domain.Song{
		Artist:      cmd.Artist,
		Title:       cmd.Title,
		Lyrics:      songInfo.Lyrics,
		ReleaseDate: songInfo.ReleaseDate,
		Link:        songInfo.Link,
	}
	song, err = h.songRepo.CreateSong(ctx, song)

	if err != nil {
		if _, ok := custom_error.IsCustom(err); ok {
			return nil, err
		}
		return nil, custom_error.NewInternalError(err, "failed to create song")
	}

	return &CreateSongResult{
		Song: song,
	}, nil
}
