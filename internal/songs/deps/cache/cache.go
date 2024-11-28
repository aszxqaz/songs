package cache

import (
	"context"
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("key not found")

type CacheManager interface {
	Set(ctx context.Context, key string, val string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
