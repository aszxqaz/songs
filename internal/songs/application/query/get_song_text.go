package query

import (
	"context"
	"encoding/json"
	"fmt"
	"songs/internal/common/decorator"
	"songs/internal/common/helpers"
	"songs/internal/common/pagination"
	"songs/internal/songs/deps/cache"
	"songs/internal/songs/deps/repository"

	"github.com/sirupsen/logrus"
)

var getSongTextPaginationConstraints = pagination.Constraints{
	MinLimit:  1,
	MaxLimit:  200,
	MaxOffset: 200,
}

var getSongTextPaginationDefaults = pagination.Defaults{
	Limit:  1,
	Offset: 0,
}

type GetSongTextQuery struct {
	ID          int
	PagParams   pagination.ParamsOptional
	CoupletSize int
}

type GetSongTextQueryResult struct {
	Couplets   []string
	Pagination Pagination
}

type GetSongTextQueryHandler decorator.QueryHandler[GetSongTextQuery, *GetSongTextQueryResult]

type getSongTextQueryHandler struct {
	songRepo     repository.SongRepository
	cacheManager cache.CacheManager
}

func NewGetSongTextQueryHandler(
	songRepo repository.SongRepository,
	cacheManager cache.CacheManager,
	logger *logrus.Entry,
) GetSongTextQueryHandler {
	return decorator.ApplyQueryDecorators(&getSongTextQueryHandler{
		songRepo,
		cacheManager,
	}, logger)
}

// Handle implements GetSongTextQueryHandler.
func (g *getSongTextQueryHandler) Handle(ctx context.Context, query GetSongTextQuery) (*GetSongTextQueryResult, error) {
	pagParams := query.PagParams.MergeDefaults(getSongTextPaginationDefaults)
	err := pagParams.CheckConstraints(getSongTextPaginationConstraints)
	if err != nil {
		return nil, err
	}

	couplets, err := g.lookUpCoupletsInCache(ctx, query.ID)

	if err == nil {
		return &GetSongTextQueryResult{
			Couplets: pagination.StripSlice(couplets, pagParams),
			Pagination: Pagination{
				Limit:  pagParams.Limit,
				Offset: pagParams.Offset,
				Total:  len(couplets),
			},
		}, nil
	}

	song, err := g.songRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	couplets = helpers.ChunkedBySeparator(song.Lyrics, "\n", query.CoupletSize)

	go g.addCoupletsToCache(ctx, song.ID, couplets)

	return &GetSongTextQueryResult{
		Couplets: pagination.StripSlice(couplets, pagParams),
		Pagination: Pagination{
			Limit:  pagParams.Limit,
			Offset: pagParams.Offset,
			Total:  len(couplets),
		},
	}, nil
}

func (g *getSongTextQueryHandler) lookUpCoupletsInCache(ctx context.Context, songID int) ([]string, error) {
	coupletsJson, err := g.cacheManager.Get(ctx, songTextKey(songID))
	if err != nil {
		return nil, err
	}
	couplets := make([]string, 0)
	json.Unmarshal([]byte(coupletsJson), &couplets)
	return couplets, nil
}

func (g *getSongTextQueryHandler) addCoupletsToCache(ctx context.Context, songID int, couplets []string) error {
	coupletsJson, _ := json.Marshal(couplets)
	err := g.cacheManager.Set(ctx, songTextKey(songID), string(coupletsJson), 0)
	return err
}

func songTextKey(songID int) string {
	return fmt.Sprintf("couplets:%d", songID)
}

// func joinCouplets(couplets []string, pagParams pagination.Params) string {
// 	textChunks := pagination.StripSlice(couplets, pagParams)
// 	return strings.Join(textChunks, "\n")
// }
