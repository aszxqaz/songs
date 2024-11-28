package song

import (
	"songs/internal/songs/domain"
	"songs/internal/songs/ports/http/contracts"
)

func NewApiSongFromDomain(s *domain.Song) *contracts.Song {
	return &contracts.Song{
		Id:          s.ID,
		Group:       s.Artist,
		Song:        s.Title,
		Text:        s.Lyrics,
		ReleaseDate: s.ReleaseDate,
		Link:        s.Link,
	}
}
