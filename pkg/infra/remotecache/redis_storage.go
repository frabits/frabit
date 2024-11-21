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
