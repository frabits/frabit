package remotecache

import (
	"context"
	"database/sql"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/registry"
)

const (
	CacheRedis = "redis"
	CacheSql   = "sql"
)

type CacheStorage interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, expire time.Duration) error
	Delete(ctx context.Context, key string) error
	Count(ctx context.Context, prefix string) (int64, error)
}

type RemoteCache struct {
	Cfg       *config.Config
	Client    CacheStorage
	MetaStore *db.MetaStore
}

func ProviderRemoteCache(cfg *config.Config, meta *db.MetaStore) *RemoteCache {
	rc := &RemoteCache{
		Cfg:       cfg,
		MetaStore: meta,
	}
	rc.Client = newStorage(cfg, meta.DB)
	return rc
}

func (rc *RemoteCache) Get(ctx context.Context, key string) ([]byte, error) {
	return rc.Client.Get(ctx, key)
}

func (rc *RemoteCache) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	return rc.Client.Set(ctx, key, value, expire)
}

func (rc *RemoteCache) Delete(ctx context.Context, key string) error {
	return rc.Client.Delete(ctx, key)
}

func (rc *RemoteCache) Count(ctx context.Context, prefix string) (int64, error) {
	return rc.Client.Count(ctx, prefix)
}

func (rc *RemoteCache) Run(ctx context.Context) error {
	bgClient, ok := rc.Client.(registry.BackgroundService)
	if ok {
		return bgClient.Run(ctx)
	}
	<-ctx.Done()
	return ctx.Err()
}

func newStorage(cfg *config.Config, db *sql.DB) CacheStorage {
	var cache CacheStorage
	switch cfg.Server.CacheType {
	case CacheRedis:
		cache = newRedisStorage(cfg)
	default:
		cache = newSqlStorage(cfg, db)
	}

	if cfg.Server.CachePrefix != "" {
		cache = &prefixCacheStorage{cache: cache, Prefix: cfg.Server.CachePrefix}
	}

	if cfg.Server.CacheEncrypted {
		cache = &encryptedCacheStorage{cache: cache}
	}

	return cache
}

type prefixCacheStorage struct {
	cache  CacheStorage
	Prefix string
}

func (pcs *prefixCacheStorage) Get(ctx context.Context, key string) ([]byte, error) {
	return pcs.cache.Get(ctx, pcs.Prefix+key)
}

func (pcs *prefixCacheStorage) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	return pcs.cache.Set(ctx, pcs.Prefix+key, value, expire)
}

func (pcs *prefixCacheStorage) Delete(ctx context.Context, key string) error {
	return pcs.cache.Delete(ctx, pcs.Prefix+key)
}

func (pcs *prefixCacheStorage) Count(ctx context.Context, prefix string) (int64, error) {
	return pcs.Count(ctx, pcs.Prefix+prefix)
}

type encryptedCacheStorage struct {
	cache CacheStorage
	enc   encryptSrv
}

type encryptSrv interface {
	Encrypt(context.Context, []byte) ([]byte, error)
	Decrypt(context.Context, []byte) ([]byte, error)
}

func (ecs *encryptedCacheStorage) Get(ctx context.Context, key string) ([]byte, error) {
	rawVal, err := ecs.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return ecs.enc.Decrypt(ctx, rawVal)
}

func (ecs *encryptedCacheStorage) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	encValue, err := ecs.enc.Encrypt(ctx, value)
	if err != nil {
		return err
	}
	return ecs.cache.Set(ctx, key, encValue, expire)
}

func (ecs *encryptedCacheStorage) Delete(ctx context.Context, key string) error {
	return ecs.cache.Delete(ctx, key)
}

func (ecs *encryptedCacheStorage) Count(ctx context.Context, prefix string) (int64, error) {
	return ecs.cache.Count(ctx, prefix)
}
