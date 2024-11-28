package main

import (
	"context"
	"io/fs"
	"net/http"
	"songs/internal/common/env"
	"songs/internal/common/logs"
	"songs/internal/common/server"
	"songs/internal/openapi"
	"songs/internal/songs/application"
	"songs/internal/songs/deps/cache/redis"
	"songs/internal/songs/deps/info"
	"songs/internal/songs/deps/repository/gorm"
	songs_ports_http "songs/internal/songs/ports/http"
	"songs/internal/songs/ports/http/contracts"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/swaggest/swgui/v5emb"
)

var (
	port = env.String("SONGS_PORT", "8080")

	gormConfig = &gorm.Config{
		Host:     env.String("POSTGRES_HOST", "localhost"),
		Port:     env.Int("POSTGRES_PORT", 5432),
		User:     env.String("POSTGRES_USER", "postgres"),
		Password: env.String("POSTGRES_PASSWORD", "postgres"),
		Database: env.String("POSTGRES_DB", "postgres"),
	}

	redisConfig = &redis.Config{
		Host:     env.String("REDIS_HOST", "localhost"),
		Port:     env.String("REDIS_PORT", "6379"),
		Password: env.String("REDIS_PASSWORD", "bJKA1oYa71cRYr51"),
	}

	songInfoServiceConfig = &info.Config{
		BaseURL: env.String("SONG_INFO_SERVICE_SCHEME", "http") + "://" +
			env.String("SONG_INFO_SERVICE_HOST", "localhost") + ":" +
			env.String("SONG_INFO_SERVICE_PORT", "3001"),
	}
)

func main() {
	logs.Init()
	ctx := context.Background()

	db := gorm.MustConnect(gormConfig)
	songRepo := gorm.NewPgGormSongRepository(db)
	songInfoService := info.NewSongInfoService(songInfoServiceConfig)
	cacheManager := redis.NewRedisListCacheManager(redisConfig)

	logger := logrus.NewEntry(logrus.StandardLogger())

	application := application.NewApplication(ctx, logger, songRepo, songInfoService, cacheManager)

	server.RunHTTPServer(port, func(router chi.Router) http.Handler {
		staticFs, err := fs.Sub(openapi.Spec, "spec")
		if err != nil {
			panic(err)
		}

		router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFs))))

		router.Handle("/*", v5emb.New(
			"Songs API",
			"http://localhost:"+port+"/static/songs.yml",
			"/",
		))

		return contracts.HandlerFromMux(
			songs_ports_http.NewHttpServer(application),
			router,
		)
	})
}
