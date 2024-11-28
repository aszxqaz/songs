package song

import (
	"net/http"
	"songs/internal/common/pagination"
	"songs/internal/common/server/httperr"
	"songs/internal/songs/application/query"
	"songs/internal/songs/ports/http/contracts"

	"github.com/go-chi/render"
)

const (
	getSongTextDefaultCoupletSize = 4
	getSongTextDefaultLimit       = 1
	coupletSize                   = 4
)

// GetSongText implements ServerInterface.
func (h Handler) GetSongText(w http.ResponseWriter, r *http.Request, songID int, params contracts.GetSongTextParams) {
	result, err := h.app.Queries.GetSongText.Handle(
		r.Context(),
		query.GetSongTextQuery{
			ID: songID,
			PagParams: pagination.ParamsOptional{
				Limit:  params.Limit,
				Offset: params.Offset,
			},
			CoupletSize: coupletSize,
		},
	)

	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	body := map[string]interface{}{
		"couplets":   result.Couplets,
		"pagination": result.Pagination,
	}

	render.JSON(w, r, body)
}
