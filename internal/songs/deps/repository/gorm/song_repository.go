package gorm

import (
	domain_song "songs/internal/songs/deps/repository"

	"gorm.io/gorm"
)

type PgGormSongRepository struct {
	db *gorm.DB
}

func NewPgGormSongRepository(db *gorm.DB) domain_song.SongRepository {
	db.AutoMigrate(
		&songModel{},
	)
	return &PgGormSongRepository{db}
}
