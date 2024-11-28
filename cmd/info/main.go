package main

import (
	"net/http"
	"songs/internal/common/env"
	"songs/internal/common/logs"
	"songs/internal/common/server"
	info_ports_http "songs/internal/info/ports/http"

	"github.com/go-chi/chi/v5"
)

var (
	port = env.String("SONG_INFO_SERVICE_PORT", "3001")
)

func main() {
	logs.Init()
	server.RunHTTPServer(port, func(router chi.Router) http.Handler {
		return info_ports_http.HandlerFromMux(
			info_ports_http.NewHttpServer(),
			router,
		)
	})
}
