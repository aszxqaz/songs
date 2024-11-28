package info

import (
	"context"
	"songs/internal/common/client/info"
	custom_error "songs/internal/common/custom_error"
)

type songInfoService struct {
	client *info.ClientWithResponses
}

type SongInfo struct {
	ReleaseDate string
	Lyrics      string
	Link        string
}

type Query struct {
	Artist string
	Title  string
}

type SongInfoAPIService interface {
	GetSongInfo(ctx context.Context, q Query) (*SongInfo, error)
}

type Config struct {
	BaseURL string
}

func NewSongInfoService(config *Config) SongInfoAPIService {
	client, err := info.NewClientWithResponses(config.BaseURL)
	if err != nil {
		panic(err)
	}
	return songInfoService{client}
}

// GetSongInfo implements SongInfoAPIService.
func (s songInfoService) GetSongInfo(ctx context.Context, q Query) (*SongInfo, error) {
	rsp, err := s.client.GetInfoWithResponse(ctx, &info.GetInfoParams{
		Group: q.Artist,
		Song:  q.Title,
	})
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode() != 200 {
		if rsp.StatusCode() == 404 {
			return nil, custom_error.NewNotFoundError(err, "details of the song not found")
		}

		return nil, err
	}

	return &SongInfo{
		ReleaseDate: rsp.JSON200.ReleaseDate,
		Lyrics:      rsp.JSON200.Text,
		Link:        rsp.JSON200.Link,
	}, nil
}
