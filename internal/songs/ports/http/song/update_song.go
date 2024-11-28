package song

import (
	"net/http"
	"songs/internal/common/custom_error"
	"songs/internal/common/server/httperr"
	"songs/internal/common/validator"
	"songs/internal/songs/application/command"
	"songs/internal/songs/ports/http/contracts"

	"github.com/go-chi/render"
)

// UpdateSong implements ServerInterface.
func (h Handler) UpdateSong(w http.ResponseWriter, r *http.Request, songID int) {
	var data contracts.UpdateSongJSONRequestBody
	if err := render.Decode(r, &data); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	if err := validator.Validate.Struct(&data); err != nil {
		httperr.RespondWithError(custom_error.NewBadInputError(err, err.Error()), w, r)
		return
	}

	result, err := h.app.Commands.UpdateSong.Handle(r.Context(), command.UpdateSongCommand{
		ID:          songID,
		Artist:      data.Group,
		Title:       data.Song,
		Lyrics:      data.Text,
		ReleaseDate: data.ReleaseDate,
		Link:        data.Link,
	})
	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	render.JSON(w, r, NewApiSongFromDomain(result.Song))
}
