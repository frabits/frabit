/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package ctl

import (
	"context"
	"database/sql"
)

type DBType string

const (
	MySQL      DBType = "MYSQL"
	Redis      DBType = "REDIS"
	MongoDB    DBType = "MONGODB"
	ClickHouse DBType = "CLICKHOUSE"
)

type DBConnInfo struct {
	// General DB filed
	Host   string
	Port   string
	User   string
	Passwd string

	// For specific DBType
	// For MySQL/ClickHouse
	Database string

	// For MongoDB
	SRV    bool
	AuthDB string
}

// Driver is the interface for supported database driver.
type Driver interface {
	Ping(ctx context.Context) error
	Open(ctx context.Context, dbType DBType, config DBConnInfo) (Driver, error)
	Close(ctx context.Context) error
	GetType() DBType
	GetDBConn(ctx context.Context) (*sql.DB, error)
	Execute(ctx context.Context, statement string, createDatabase bool) (int64, error)
	QueryConn(ctx context.Context, conn *sql.Conn, statement string) ([]interface{}, error)
}
