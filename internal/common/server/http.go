package server

import (
	"net/http"
	"songs/internal/common/logs"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(port string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	// apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.NoCache)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/", createHandler(apiRouter))

	logrus.Infof("HTTP server listening on port %s", port)

	err := http.ListenAndServe(":"+port, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}
