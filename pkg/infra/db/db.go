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
	dc "github.com/frabits/frabit/pkg/common/db_connector"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MetaStore struct {
	cfg *config.Config
	log *slog.Logger
	DB  *sql.DB
	Gdb *gorm.DB
}

var metaStore *gorm.DB

func New(conf *config.Config) (*MetaStore, error) {
	ms := &MetaStore{
		cfg: conf,
		log: log.New("meta.store"),
	}
	dbConnectConfig := dc.MysqlConnectionCfg{
		Username:  conf.DB.UserName,
		Password:  conf.DB.Password,
		Hostname:  conf.DB.Host,
		Port:      strconv.Itoa(int(conf.DB.Port)),
		DefaultDb: conf.DB.Database,
		Params:    []string{"timeout=5s", "charset=utf8mb4", "parseTime=True", "loc=Local", "readTimeout=30s"},
	}

	dsn := dbConnectConfig.String()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		ms.log.Error("backend metaStore not ready", "Error", err.Error())
		return ms, err
	}
	ms.Gdb = db
	stdDb, err := db.DB()
	ms.DB = stdDb
	stdDb.SetConnMaxLifetime(10 * time.Minute)
	stdDb.SetMaxIdleConns(50)
	stdDb.SetMaxOpenConns(10)

	return ms, nil
}

func DB() *gorm.DB {
	return metaStore
}

// OpenFrabit returns the DB instance for the frabit backed database
func (ms *MetaStore) OpenFrabit() (db *sql.DB, err error) {
	// first time ever we talk to MySQL
	query := fmt.Sprintf("create database if not exists %s", ms.cfg.DB.Database)
	if _, err := ms.DB.Exec(query); err != nil {
		return db, err
	}
	if !ms.cfg.DB.SkipDatabaseUpdate && !ms.alreadyDeployed() {
		ms.log.Info("Start initiate backend metaStore")
		err := ms.initFrabitDB(ms.DB)
		if err != nil {
			ms.log.Error("Initiate backend metaStore failed", "Error", err.Error())
			return nil, err
		}
		ms.log.Info("Successfully create all tables")
	}
	maxIdleConns := int(100)
	if maxIdleConns < 10 {
		ms.log.Warn("MaxIdleConnections less than 10,already reset to 10")
		maxIdleConns = 10
	}
	ms.log.Info("Connecting to backend metastore", "host", ms.cfg.DB.Host, "port", ms.cfg.DB.Port, "maxConnections", ms.cfg.DB.MaxPoolConnections)
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
	license := fmt.Sprintf(`insert into license(license_level,current_license,last_license,created_at,updated_at) values("%s","%s","%s","%s","%s")`, constant.Community, "", "", initDatetime, initDatetime)

	initPasswd := utils.GenRandom(32)
	rands := utils.GenRandom(12)
	hashPassword := utils.GeneratePassword(initPasswd)
	ms.log.Info("InitPassword generated", "Username", "admin", "Password", initPasswd)
	adminAccount := fmt.Sprintf(`insert into user(login,username,email,password,rands,is_admin,is_disabled,is_external,is_email_verified,theme,org_id,created_at,updated_at,last_seen_at) values("%s","%s","%s","%s","%s",%d,%d,%d,%d,"%s",%d,"%s","%s","%s")`, constant.ADMIN, constant.ADMIN, "admin@frabit.com", hashPassword, rands, 1, 0, 0, 1, constant.Dark, 1, initDatetime, initDatetime, initDatetime)
	org := fmt.Sprintf(`insert into org(name,description,created_at,updated_at) values("%s","%s","%s","%s")`, constant.MainOrg, "Default org", initDatetime, initDatetime)
	version := "v2.2.2"
	initVersion := fmt.Sprintf(`insert into version(version,created_at,updated_at) values("%s","%s","%s")`, version, initDatetime, initDatetime)
	initData = append(initData, license, adminAccount, org, initVersion)
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
			return true
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
