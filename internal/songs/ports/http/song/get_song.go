package song

import (
	"net/http"
	"songs/internal/common/server/httperr"
	"songs/internal/songs/application/query"

	"github.com/go-chi/render"
)

// GetSong implements ServerInterface.
func (h Handler) GetSong(w http.ResponseWriter, r *http.Request, songID int) {
	queryResult, err := h.app.Queries.GetSong.Handle(r.Context(), query.GetSongQuery{
		ID: songID,
	})

	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	render.JSON(w, r, NewApiSongFromDomain(queryResult.Song))
}
