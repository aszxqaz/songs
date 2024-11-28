package song

import (
	"net/http"
	"songs/internal/common/server/httperr"
	"songs/internal/common/validator"
	"songs/internal/songs/application/command"
	"songs/internal/songs/ports/http/contracts"

	"github.com/go-chi/render"
)

// CreateSong implements ServerInterface.
func (h Handler) CreateSong(w http.ResponseWriter, r *http.Request) {
	var data contracts.CreateSongJSONRequestBody
	if err := render.Decode(r, &data); err != nil {
		httperr.BadRequest(err, "Malformed request body", w, r)
		return
	}

	if err := validator.Validate.Struct(&data); err != nil {
		httperr.BadRequest(err, err.Error(), w, r)
		return
	}

	result, err := h.app.Commands.CreateSong.Handle(r.Context(), command.CreateSongCommand{
		Artist: data.Group,
		Title:  data.Song,
	})
	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	song := NewApiSongFromDomain(result.Song)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, song)
}
