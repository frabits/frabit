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
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/frabits/frabit/pkg/common/constant"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"

	"github.com/percona/go-mysql/dsn"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MetaStore struct {
	cfg *config.Config
	log *slog.Logger
	DB  *sql.DB
}

var DB *gorm.DB

func New(conf *config.Config) (*MetaStore, error) {
	ms := &MetaStore{
		cfg: conf,
		log: log.New("meta.store"),
	}
	dsn := dsn.DSN{
		Username:  conf.DB.UserName,
		Password:  conf.DB.Password,
		Hostname:  conf.DB.Host,
		Port:      strconv.Itoa(int(conf.DB.Port)),
		DefaultDb: conf.DB.Database,
	}
	db, err := gorm.Open(mysql.Open(dsn.String()), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return ms, err
	}
	DB = db
	stdDb, err := db.DB()
	ms.DB = stdDb

	return ms, nil
}

// OpenFrabit returns the DB instance for the frabit backed database
func (ms *MetaStore) OpenFrabit() (db *sql.DB, err error) {
	// first time ever we talk to MySQL
	query := fmt.Sprintf("create database if not exists %s", config.Conf.DB.Database)
	if _, err := ms.DB.Exec(query); err != nil {
		return db, err
	}
	if !config.Conf.DB.SkipDatabaseUpdate && !ms.alreadyDeployed() {
		err := ms.initFrabitDB(ms.DB)
		if err != nil {
			return nil, err
		}
	}
	maxIdleConns := int(100)
	if maxIdleConns < 10 {
		maxIdleConns = 10
	}
	ms.log.Info("Connecting to backend metastore", "host", config.Conf.DB.Host, "port", config.Conf.DB.Port, "maxConnections", config.Conf.DB.MaxPoolConnections)
	ms.DB.SetMaxIdleConns(maxIdleConns)
	return ms.DB, err
}

// initFrabitDB attempts to create/upgrade the frabit backend database. It is created once in the
// application's lifetime.
func (ms *MetaStore) initFrabitDB(db *sql.DB) error {
	if err := ms.deployStatements(db, generateSQLBase); err != nil {
		return err
	}
	if err := ms.deployStatements(db, generateSQLPatches); err != nil {
		return err
	}
	initSQL := ms.genInitialData()
	if err := ms.deployStatements(db, initSQL); err != nil {
		return err
	}
	return nil
}

func (ms *MetaStore) genInitialData() []string {
	initDatetime := time.Now().Format(time.DateTime)
	initData := make([]string, 0)
	license := fmt.Sprintf(`insert into license(license_level,current_license,last_license,create_at,update_at) values("%s","%s","%s","%s","%s")`, constant.Community, "", "", initDatetime, initDatetime)

	initPasswd := utils.GenRandom(32)
	hashPassword := utils.GeneratePassword(initPasswd)
	ms.log.Info("InitPassword generated", "Username", "admin", "Password", initPasswd, "HashPassword", hashPassword)
	adminAccount := fmt.Sprintf(`insert into user(username,email,password,is_disabled) values("%s","%s","%s",%d)`, constant.ADMIN, "admin@frabit.com", hashPassword, 0)
	version := "v2.2.2"
	initVersion := fmt.Sprintf(`insert into version(version,create_at,update_at) values("%s","%s","%s")`, version, initDatetime, initDatetime)
	initData = append(initData, license, adminAccount, initVersion)
	return initData
}

// alreadyDeployed check tables
func (ms *MetaStore) alreadyDeployed() bool {
	query := fmt.Sprintf(`select count(*) table_num from information_schema.tables where table_schema="%s";`, ms.cfg.DB.Database)
	rows, err := ms.DB.Query(query)
	defer rows.Close()
	if err != nil {
		ms.log.Error("can not detect table number", "Error", err.Error())
		return true
	}
	for rows.Next() {
		var tableNum uint
		if err := rows.Scan(&tableNum); err != nil {
			ms.log.Error("can not detect table num", "Error", err.Error())
			return true
		}
		if tableNum > 0 {
			ms.log.Info("table already created")
			return false
		}
	}
	return false
}

// deployStatements will issue given sql queries that are not already known to be deployed.
// This iterates both lists (to-run and already-deployed) and also verifies no contradictions.
func (ms *MetaStore) deployStatements(db *sql.DB, queries []string) error {
	tx, err := db.Begin()
	if err != nil {
		ms.log.Error("Start transaction failed", "Error", err.Error())
	}
	for i, query := range queries {
		if i == 0 {
			//log.Debugf("sql_mode is: %+v", originalSqlMode)
		}
		if _, err := tx.Exec(query); err != nil {
			if strings.Contains(err.Error(), "syntax error") {
				return err
			}
			if !strings.Contains(err.Error(), "duplicate column name") &&
				!strings.Contains(err.Error(), "Duplicate column name") &&
				!strings.Contains(err.Error(), "check that column/key exists") &&
				!strings.Contains(err.Error(), "already exists") &&
				!strings.Contains(err.Error(), "Duplicate key name") {
				ms.log.Error("Error initiating frabit", "Error", err.Error(), "query", query)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		ms.log.Error("Commit transaction failed", "Error", err.Error())
	}
	return nil
}
