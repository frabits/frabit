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

package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/operator"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// OperatorClickHouse implement a ClickHouse DBOperator
type OperatorClickHouse struct {
	log        *slog.Logger
	connConfig operator.DBConnInfo
	dbType     operator.DBType
	db         *sql.DB
}

func (op *OperatorClickHouse) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.DBOperator, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	op.log.Info("Connect to Clickhouse", "host", addr)
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
	op.dbType = dbName
	op.db = conn

	return op, nil
}

func (op *OperatorClickHouse) GetType() operator.DBType {
	return operator.ClickHouse
}

func (op *OperatorClickHouse) Ping(ctx context.Context) error {
	return op.db.PingContext(ctx)
}

func (op *OperatorClickHouse) Close(ctx context.Context) error {
	return op.db.Close()
}
