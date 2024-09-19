package remotecache

import (
	"context"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client redis.UniversalClient
}

func parseRedisOpts(redisConn string) *redis.UniversalOptions {
	opts := &redis.UniversalOptions{}

	return opts
}

func newRedisStorage(cfg *config.Config) *RedisStorage {
	opts := parseRedisOpts(cfg.Server.CacheConnection)
	redisClient := redis.NewUniversalClient(opts)
	return &RedisStorage{client: redisClient}
}

func (rs *RedisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	ret := rs.client.Get(ctx, key)
	return ret.Bytes()
}

func (rs *RedisStorage) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	ret := rs.client.Set(ctx, key, value, expire)
	return ret.Err()
}

func (rs *RedisStorage) Delete(ctx context.Context, key string) error {
	ret := rs.client.Del(ctx, key)
	return ret.Err()
}

func (rs *RedisStorage) Count(ctx context.Context, prefix string) (int64, error) {
	ret := rs.client.Keys(ctx, prefix)
	return int64(len(ret.Val())), ret.Err()
}
