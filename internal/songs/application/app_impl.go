package application

import (
	"context"
	"songs/internal/songs/application/command"
	"songs/internal/songs/application/query"
	"songs/internal/songs/deps/cache"
	"songs/internal/songs/deps/info"
	"songs/internal/songs/deps/repository"

	"github.com/sirupsen/logrus"
)

func NewApplication(
	ctx context.Context,
	logger *logrus.Entry,
	songRepo repository.SongRepository,
	songInfoService info.SongInfoAPIService,
	cacheManager cache.CacheManager,
) Application {
	return Application{
		Commands: Commands{
			CreateSong: command.NewCreateSongCommandHandler(songRepo, songInfoService, logger),
			DeleteSong: command.NewDeleteSongCommandHandler(songRepo, logger),
			UpdateSong: command.NewUpdateSongCommandHandler(songRepo, logger),
		},
		Queries: Queries{
			GetSongs:    query.NewGetSongsQueryHandler(songRepo, logger),
			GetSong:     query.NewGetSongQueryHandler(songRepo, logger),
			GetSongText: query.NewGetSongTextQueryHandler(songRepo, cacheManager, logger),
		},
	}
}
