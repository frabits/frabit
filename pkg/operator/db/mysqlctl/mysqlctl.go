// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
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

package mysqlctl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/frabits/frabit/pkg/operator"
)

type Operator struct {
	Host   string
	Port   uint32
	Passwd string
	DBName operator.DBType
	db     *sql.DB
}

func (op *Operator) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.DBOperator, error) {
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
	op.DBName = dbName
	op.db = db

	return op, nil
}

func (op *Operator) GetType() operator.DBType {
	return operator.MySQL
}

func (op *Operator) Ping(ctx context.Context) error {
	return op.db.PingContext(ctx)
}

func (op *Operator) Close(ctx context.Context) error {
	return op.db.Close()
}
