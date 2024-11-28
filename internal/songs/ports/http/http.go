package http

import (
	"songs/internal/songs/application"
	"songs/internal/songs/ports/http/contracts"
	"songs/internal/songs/ports/http/song"
)

type HttpServer struct {
	song.Handler
}

func NewHttpServer(app application.Application) contracts.ServerInterface {
	return HttpServer{
		Handler: song.NewHandler(&app),
	}
}
