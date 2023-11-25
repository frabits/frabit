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

package meta_store

import (
	"database/sql"
	"fmt"
	"github.com/frabits/frabit/pkg/common/log"
	"github.com/frabits/frabit/pkg/config"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// OpenFrabit returns the DB instance for the frabit backed database
func OpenFrabit() (db *sql.DB, err error) {
	// first time ever we talk to MySQL
	query := fmt.Sprintf("create database if not exists %s", config.Conf.DB.MySQLFrabitDatabase)
	if _, err := db.Exec(query); err != nil {
		return db, err
	}
	if !config.Conf.DB.SkipFrabitDatabaseUpdate {
		err := initFrabitDB(db)
		if err != nil {
			return nil, err
		}
	}
	maxIdleConns := int(100)
	if maxIdleConns < 10 {
		maxIdleConns = 10
	}
	log.Info("Connecting to backend  %s:%d: maxConnections: %d, maxIdleConns: %d",
		zap.String("host", config.Conf.DB.MySQLFrabitHost),
		zap.String("port", config.Conf.DB.MySQLFrabitPort),
		zap.Int("maxConnections", config.Conf.DB.MySQLFrabitMaxPoolConnections),
		zap.Int("maxIdleConns", maxIdleConns))
	db.SetMaxIdleConns(maxIdleConns)
	return db, err
}

// initFrabitDB attempts to create/upgrade the frabit backend database. It is created once in the
// application's lifetime.
func initFrabitDB(db *sql.DB) error {
	if err := deployStatements(db, generateSQLBase); err != nil {
		return err
	}
	if err := deployStatements(db, generateSQLPatches); err != nil {
		return err
	}
	return nil
}

// deployStatements will issue given sql queries that are not already known to be deployed.
// This iterates both lists (to-run and already-deployed) and also verifies no contradictions.
func deployStatements(db *sql.DB, queries []string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Error("Start transaction failed", zap.Error(err))
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
				log.Error("Error initiating frabit", zap.Error(err), zap.String("query", query))
			}
		}
	}
	if err := tx.Commit(); err != nil {
		log.Error("Commit transaction failed", zap.Error(err))
	}
	return nil
}
