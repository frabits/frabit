// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
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

package redisctl

import (
	"context"
	"fmt"

	"github.com/frabits/frabit/pkg/operator"

	"github.com/redis/go-redis/v9"
)

type RedisMode string

const (
	Standalone RedisMode = "STANDALONE"
	Sentinel   RedisMode = "SENTINEL"
	Cluster    RedisMode = "CLUSTER"
)

type Operator struct {
	Host   string
	Port   uint32
	Passwd string
	DBName operator.DBType
	Mode   RedisMode
	Client *redis.Client
}

func (op *Operator) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.DBOperator, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Passwd, // no password set
		DB:       0,             // use default DB
	})
	op.Client = redis
	op.DBName = dbName
	return op, nil
}

func (op *Operator) Ping(ctx context.Context) error {
	_ = op.Client.Ping(ctx)
	return nil
}

func (op *Operator) Close(ctx context.Context) error {
	return op.Client.Close()
}
func (op *Operator) GetType() operator.DBType {
	return operator.Redis
}
