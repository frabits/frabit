package remotecache

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
)

type SqlStorage struct {
	log *slog.Logger
	db  *sql.DB
}

func newSqlStorage(cfg *config.Config, db *sql.DB) *SqlStorage {
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
	query := fmt.Sprintf(`insert into data_cache(data_key,data_value,create_at,expire) values("%s","%s","%v","%v");`, key, encoded, now, expireDatetime)
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
	query := fmt.Sprintf(`delete from data_cache where expire < "%v" and expire is not null;`, now)
	ss.db.ExecContext(ctx, query)
}

type DataCache struct {
	DataKey   string    `json:"data_key,omitempty"`
	DataValue string    `json:"data_value,omitempty"`
	CreateAt  time.Time `json:"create_at"`
	Expire    time.Time `json:"expire,omitempty"`
}
