package song

import (
	"net/http"
	"songs/internal/common/server/httperr"
	"songs/internal/songs/application/command"
)

// DeleteSong implements ServerInterface.
func (h Handler) DeleteSong(w http.ResponseWriter, r *http.Request, songID int) {
	_, err := h.app.Commands.DeleteSong.Handle(r.Context(), command.DeleteSongCommand{ID: songID})
	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
