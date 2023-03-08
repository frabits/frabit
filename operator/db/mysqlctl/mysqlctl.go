/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package mysqlctl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/frabits/frabit/operator"
)

type Driver struct {
	Host   string
	Port   uint32
	Passwd string
	DBName operator.DBType
	db     *sql.DB
}

func (driver *Driver) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.Driver, error) {
	protocol := "tcp"
	if strings.HasPrefix(config.Host, "/") {
		protocol = "unix"
	}

	params := []string{"multiStatements=true"}

	port := config.Port
	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf("%s@%s(%s:%s)/%s?%s", config.User, protocol, config.Host, port, config.Database, strings.Join(params, "&"))
	if config.Passwd != "" {
		dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", config.User, config.Passwd, protocol, config.Host, port, config.Database, strings.Join(params, "&"))
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	driver.DBName = dbName
	driver.db = db

	return driver, nil
}

func (driver *Driver) GetType() operator.DBType {
	return operator.MySQL
}

func (driver *Driver) Ping(ctx context.Context) error {
	return driver.db.PingContext(ctx)
}
