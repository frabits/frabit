/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package ctl

import "context"

type DBType string

const (
	Redis   DBType = "REDIS"
	MongoDB DBType = "MONGODB"
)

type DBConnInfo struct {
	Host   string
	Port   uint32
	User   string
	Passwd string
}
type Driver interface {
	Ping(ctx context.Context, db DBType, connInfo DBConnInfo)
}
