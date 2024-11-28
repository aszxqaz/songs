package gorm

import (
	"songs/internal/common/date"
	"songs/internal/songs/domain"
	"time"

	"gorm.io/gorm"
)

type songModel struct {
	gorm.Model
	Artist      string `gorm:"uniqueIndex:idx_artist_title"`
	Title       string `gorm:"uniqueIndex:idx_artist_title"`
	Lyrics      string
	ReleaseDate time.Time
	Link        string
}

func (s *songModel) toDomain() *domain.Song {
	return &domain.Song{
		ID:          int(s.ID),
		Artist:      s.Artist,
		Title:       s.Title,
		Lyrics:      s.Lyrics,
		ReleaseDate: date.TimeToDMY(s.ReleaseDate),
		Link:        s.Link,
	}
}

func (s *songModel) mergeWithDomain(d *domain.Song) {
	s.Artist = d.Artist
	s.Title = d.Title
	s.Lyrics = d.Lyrics
	s.ReleaseDate = date.DmyToTime(d.ReleaseDate)
	s.Link = d.Link
}

func newSongModelFromDomain(domain *domain.Song) *songModel {
	return &songModel{
		Artist:      domain.Artist,
		Title:       domain.Title,
		Lyrics:      domain.Lyrics,
		ReleaseDate: date.DmyToTime(domain.ReleaseDate),
		Link:        domain.Link,
	}
}
