/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package redisctl

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/frabits/frabit/ctl"
)

type RedisMode string

const (
	Standalone RedisMode = "STANDALONE"
	Sentinel   RedisMode = "SENTINEL"
	Cluster    RedisMode = "CLUSTER"
)

type Driver struct {
	Host   string
	Port   uint32
	Passwd string
	DBName ctl.DBType
	Mode   RedisMode
	Client *redis.Client
}

func (driver *Driver) Open(ctx context.Context, dbName ctl.DBType, config ctl.DBConnInfo) (ctl.Driver, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Passwd, // no password set
		DB:       0,             // use default DB
	})
	driver.Client = redis
	driver.DBName = dbName
	return driver, nil
}

func (driver *Driver) Ping(ctx context.Context) error {
	return nil
}

func (driver *Driver) GetType() ctl.DBType {
	return ctl.Redis
}
