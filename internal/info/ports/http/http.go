package http

import (
	"net/http"
	"songs/internal/info/fixtures"

	"github.com/go-chi/render"
)

// Mock SongDetail service that provides information
// about a requested song
type HttpServer struct{}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

// GetInfo implements ServerInterface.
func (h *HttpServer) GetInfo(w http.ResponseWriter, r *http.Request, params GetInfoParams) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, SongDetail{
		ReleaseDate: fixtures.Date,
		Link:        fixtures.Link,
		Text:        fixtures.Lyrics,
	})
}
