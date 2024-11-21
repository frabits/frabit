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

package remotecache

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/infra/log"
)

type SqlStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func newSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{
		log: log.New("remotecache.sqlStorage"),
		db:  db,
	}
}

func (ss *SqlStorage) Get(ctx context.Context, key string) ([]byte, error) {
	var data DataCache
	query := fmt.Sprintf(`select data_value from data_cache where data_key = "%s";`, key)
	results, err := ss.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		if err := results.Scan(&data.DataValue); err != nil {
			return nil, err
		}
	}
	return base64.StdEncoding.DecodeString(data.DataValue)
}

func (ss *SqlStorage) Set(ctx context.Context, key string, value []byte, expire time.Duration) error {
	encoded := base64.StdEncoding.EncodeToString(value)
	now := time.Now()
	expireDatetime := now.Add(expire)
	ss.log.Info("create_at", now)
	query := fmt.Sprintf(`insert into data_cache(data_key,data_value,created_at,expired_at) values("%s","%s","%v","%v");`, key, encoded, now.Format(time.DateTime), expireDatetime.Format(time.DateTime))
	_, err := ss.db.ExecContext(ctx, query)
	return err
}

func (ss *SqlStorage) Delete(ctx context.Context, key string) error {
	query := fmt.Sprintf(`delete from data_cache where data_key = "%s";`, key)
	_, err := ss.db.ExecContext(ctx, query)
	return err
}

func (ss *SqlStorage) Count(ctx context.Context, prefix string) (int64, error) {
	var keyNum int64
	query := fmt.Sprintf(`select count(*) as keyNum from data_cache where data_key like "%s";`, prefix)
	results, err := ss.db.QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}
	for results.Next() {
		if err := results.Scan(&keyNum); err != nil {
			return 0, err
		}
	}
	return keyNum, nil
}

func (ss *SqlStorage) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Minute * 10)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			ss.log.Info("Start purge expired remoteCache")
			ss.purgeExpireKey(ctx)
		}
	}
}

func (ss *SqlStorage) purgeExpireKey(ctx context.Context) {
	now := time.Now()
	query := fmt.Sprintf(`delete from data_cache where expired_at < "%v" and expired_at is not null;`, now.Format(time.DateTime))
	if result, err := ss.db.ExecContext(ctx, query); err != nil {
		ss.log.Error("purge expire key failed", "Error", err.Error())
	} else {
		purgeNum, _ := result.RowsAffected()
		ss.log.Info("purge expire key successfully", "Expired number", purgeNum)
	}
}

type DataCache struct {
	DataKey   string    `json:"data_key,omitempty"`
	DataValue string    `json:"data_value,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}
