package httperr

import (
	"log/slog"
	"net/http"
	custom_error "songs/internal/common/custom_error"
	"songs/internal/common/logs"

	"github.com/go-chi/render"
)

func InternalError(err error, message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(err, message, w, r, http.StatusInternalServerError)
}

func BadRequest(err error, message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(err, message, w, r, http.StatusBadRequest)
}

func NotFound(err error, message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(err, message, w, r, http.StatusNotFound)
}

func Conflict(err error, message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(err, message, w, r, http.StatusConflict)
}

func RespondWithError(err error, w http.ResponseWriter, r *http.Request) {
	customError, ok := err.(*custom_error.Error)
	if !ok {
		slog.Info("HERE")
		InternalError(err, "something went wrong", w, r)
		return
	}

	switch customError.Code() {
	case custom_error.ErrInvalidInput:
		BadRequest(customError, customError.Message(), w, r)
	case custom_error.ErrNotFound:
		NotFound(customError, customError.Message(), w, r)
	case custom_error.ErrConflict:
		Conflict(customError, customError.Message(), w, r)
	default:
		InternalError(customError, customError.Message(), w, r)
	}
}

func respondWithError(err error, message string, w http.ResponseWriter, r *http.Request, status int) {
	if err != nil {
		logs.GetLogEntry(r).WithError(err).WithField("message", message).Warn("respond with error")
	}

	render.Status(r, status)
	render.JSON(w, r, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    status,
			"message": message,
		},
	})
}

type ErrorResponse struct {
	Message    string `json:"message"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}
