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

package db_connector

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/frabits/frabit/pkg/common/constant"
	"strings"
)

type MysqlConnectionCfg struct {
	Username string
	Password string
	Hostname string
	Port     string
	Socket   string
	//
	DefaultsFile string
	Protocol     string
	//
	DefaultDb string
	Params    []string
}

func (mcc MysqlConnectionCfg) String() string {
	dsnString := ""

	// Socket takes priority if set and protocol isn't tcp.
	if mcc.Socket != "" && mcc.Protocol != "tcp" {
		dsnString = fmt.Sprintf("%s:%s@unix(%s)",
			mcc.Username,
			mcc.Password,
			mcc.Socket,
		)
	} else {
		if mcc.Hostname == "" {
			mcc.Hostname = "localhost"
		}
		if mcc.Port == "" {
			mcc.Port = "3306"
		}
		dsnString = fmt.Sprintf("%s:%s@tcp(%s:%s)",
			mcc.Username,
			mcc.Password,
			mcc.Hostname,
			mcc.Port,
		)
	}

	dsnString += "/" + mcc.DefaultDb

	params := strings.Join(mcc.Params, "&")
	if params != "" {
		dsnString += "?" + params
	}

	return dsnString
}

type MySQL struct {
	cfg *MysqlConnectionCfg
}

func NewMySQL(cfg *MysqlConnectionCfg) *MySQL {
	return &MySQL{
		cfg: cfg,
	}
}

func (m *MySQL) Connect(ctx context.Context) *sql.DB {
	return nil
}

func (m *MySQL) DbType() constant.DbType {
	return constant.Mysql
}
