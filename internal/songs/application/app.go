package application

import (
	"songs/internal/songs/application/command"
	"songs/internal/songs/application/query"
)

type Application struct {
	Commands
	Queries
}

type Commands struct {
	CreateSong command.CreateSongCommandHandler
	DeleteSong command.DeleteSongCommandHandler
	UpdateSong command.UpdateSongCommandHandler
}

type Queries struct {
	GetSongs    query.GetSongsQueryHandler
	GetSong     query.GetSongQueryHandler
	GetSongText query.GetSongTextQueryHandler
}
