package query

import (
	"context"
	"songs/internal/common/decorator"
	"songs/internal/songs/deps/info"

	"github.com/sirupsen/logrus"
)

type GetSongInfoQuery struct {
	Artist string
	Title  string
}

type GetSongInfoQueryResult info.SongInfo

type GetSongInfoQueryHandler decorator.QueryHandler[GetSongInfoQuery, GetSongInfoQueryResult]

type getSongInfoQueryHandler struct {
	songInfoService info.SongInfoAPIService
}

func NewGetSongInfoQueryHandler(
	songInfoService info.SongInfoAPIService,
	logger *logrus.Entry,
) GetSongInfoQueryHandler {
	return decorator.ApplyQueryDecorators(&getSongInfoQueryHandler{
		songInfoService,
	}, logger)
}

// Handle implements GetSongInfoQueryHandler.
func (g *getSongInfoQueryHandler) Handle(ctx context.Context, query GetSongInfoQuery) (GetSongInfoQueryResult, error) {
	songInfo, err := g.songInfoService.GetSongInfo(ctx, info.Query{
		Artist: query.Artist,
		Title:  query.Title,
	})

	if err != nil {
		return GetSongInfoQueryResult{}, err
	}

	return GetSongInfoQueryResult{
		ReleaseDate: songInfo.ReleaseDate,
		Lyrics:      songInfo.Lyrics,
		Link:        songInfo.Link,
	}, nil
}
