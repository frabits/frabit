/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package clickhousectl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"

	"github.com/frabits/frabit/ctl"
)

type Driver struct {
	connConfig ctl.DBConnInfo
	dbType     ctl.DBType
	db         *sql.DB
}

func (driver *Driver) Open(ctx context.Context, dbName ctl.DBType, config ctl.DBConnInfo) (*Driver, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: config.Database,
			Username: config.User,
			Password: config.Passwd,
		},
		Settings: clickhouse.Settings{
			// Use a relative long value to avoid timeout on resource-intenstive query. Example failure:
			// failed: code: 160, message: Estimated query execution time (xxx seconds) is too long. Maximum: yyy. Estimated rows to process: zzzzzzzzz
			"max_execution_time": 300,
		},
		DialTimeout: 10 * time.Second,
	})
	driver.dbType = dbName
	driver.db = conn

	return driver, nil
}
