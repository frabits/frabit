/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package mysqlctl

import (
	"context"
	"database/sql"

	"github.com/frabits/frabit/ctl"
)

type Driver struct {
	Host   string
	Port   uint32
	Passwd string
	DBName ctl.DBType
	DB     *sql.DB
}

func (driver *Driver) Open(ctx context.Context, dbName ctl.DBType, config ctl.DBConnInfo) {

}
