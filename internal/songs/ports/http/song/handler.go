package song

import "songs/internal/songs/application"

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) Handler {
	return Handler{app: app}
}
