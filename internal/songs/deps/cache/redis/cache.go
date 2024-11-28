package redis

import (
	"context"
	"songs/internal/songs/deps/cache"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheManager struct {
	client *redis.Client
}

type Config struct {
	Host     string
	Port     string
	Password string
}

func NewRedisListCacheManager(config *Config) cache.CacheManager {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
	})
	return &cacheManager{client}
}

// Get implements cache.CacheManager.
func (m *cacheManager) Get(ctx context.Context, key string) (string, error) {
	val, err := m.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", cache.ErrKeyNotFound
		}
		return "", err
	}

	return val, nil
}

// Set implements cache.CacheManager.
func (m *cacheManager) Set(ctx context.Context, key string, val string, ttl time.Duration) error {
	err := m.client.Set(ctx, key, val, ttl).Err()
	return err
}
