// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remotecache

import (
	"context"
	"database/sql"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/registry"
	"github.com/frabits/frabit/pkg/services/secrets"
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
	log           *slog.Logger
	Cfg           *config.Config
	Client        CacheStorage
	MetaStore     *db.MetaStore
	secretService secrets.Service
}

func ProviderRemoteCache(cfg *config.Config, meta *db.MetaStore, secret secrets.Service) *RemoteCache {
	rc := &RemoteCache{
		log:           log.New("remoteCache"),
		Cfg:           cfg,
		MetaStore:     meta,
		secretService: secret,
	}
	rc.Client = newStorage(cfg, meta.DB, secret)
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
	cacheStorage := rc.Client
	if encCache, ok := cacheStorage.(*encryptedCacheStorage); ok {
		cacheStorage = encCache.cache
	}
	if prefixCache, ok := cacheStorage.(*prefixCacheStorage); ok {
		cacheStorage = prefixCache.cache
	}

	bgClient, ok := cacheStorage.(registry.BackgroundService)
	if ok {
		rc.log.Info("start remoteCache internalGC")
		return bgClient.Run(ctx)
	}
	<-ctx.Done()
	return ctx.Err()
}

func newStorage(cfg *config.Config, db *sql.DB, secret secrets.Service) CacheStorage {
	var cache CacheStorage
	switch cfg.Server.CacheType {
	case CacheRedis:
		cache = newRedisStorage(cfg)
	default:
		cache = newSqlStorage(db)
	}

	if cfg.Server.CachePrefix != "" {
		cache = &prefixCacheStorage{cache: cache, Prefix: cfg.Server.CachePrefix}
	}

	if cfg.Server.CacheEncrypted {
		cache = &encryptedCacheStorage{cache: cache, enc: secret}
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
	enc   secrets.Service
}

func (ecs *encryptedCacheStorage) Get(ctx context.Context, key string) ([]byte, error) {
	rawVal, err := ecs.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return ecs.enc.Decrypt(rawVal)
}

func (ecs *encryptedCacheStorage) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	encValue, err := ecs.enc.Encrypt(value)
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
