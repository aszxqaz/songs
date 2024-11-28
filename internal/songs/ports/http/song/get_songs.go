package song

import (
	"net/http"
	"songs/internal/common/date"
	"songs/internal/common/helpers"
	"songs/internal/common/pagination"
	"songs/internal/common/server/httperr"
	"songs/internal/songs/application/query"
	"songs/internal/songs/deps/repository"
	"songs/internal/songs/ports/http/contracts"

	"github.com/go-chi/render"
)

// GetSongs implements ServerInterface.
func (h Handler) GetSongs(w http.ResponseWriter, r *http.Request, params contracts.GetSongsParams) {
	searchParams := repository.FindSearchParams{
		Group:  helpers.DerefOrDefault(params.Group),
		Song:   helpers.DerefOrDefault(params.Song),
		Text:   helpers.DerefOrDefault(params.Text),
		Link:   helpers.DerefOrDefault(params.Link),
		Before: date.DmyToTime(helpers.DerefOrDefault(params.Before, date.DmyNow())),
		After:  date.DmyToTime(helpers.DerefOrDefault(params.After, "01.01.1900")),
	}

	if searchParams.After.After(searchParams.Before) {
		httperr.BadRequest(nil, "Before date should be earlier than After date", w, r)
		return
	}

	qResult, err := h.app.Queries.GetSongs.Handle(r.Context(), query.GetSongsQuery{
		SearchParams: searchParams,
		PagParams: pagination.ParamsOptional{
			Limit:  params.Limit,
			Offset: params.Offset,
		},
	})

	if err != nil {
		httperr.RespondWithError(err, w, r)
		return
	}

	songs := helpers.Map(qResult.Songs, NewApiSongFromDomain)

	render.JSON(w, r, map[string]any{
		"songs": songs,
		"pagination": contracts.Pagination{
			Limit:      qResult.Pagination.Limit,
			Offset:     qResult.Pagination.Offset,
			TotalCount: qResult.Pagination.Total,
		},
	})
}
