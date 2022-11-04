/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/frabit-io/frabit/common/log"
	"github.com/frabit-io/frabit/server/config"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// OpenFrabit returns the DB instance for the frabit backed database
func OpenFrabit() (db *sql.DB, err error) {
	// first time ever we talk to MySQL
	query := fmt.Sprintf("create database if not exists %s", config.Conf.MySQLFrabitDatabase)
	if _, err := db.Exec(query); err != nil {
		return db, err
	}
	if !config.Conf.SkipFrabitDatabaseUpdate {
		initFrabitDB(db)
	}
	maxIdleConns := int(100)
	if maxIdleConns < 10 {
		maxIdleConns = 10
	}
	log.Info("Connecting to backend  %s:%d: maxConnections: %d, maxIdleConns: %d",
		zap.String("host", config.Conf.MySQLFrabitHost),
		zap.String("port", config.Conf.MySQLFrabitPort),
		zap.Int("maxConnections", config.Conf.MySQLFrabitMaxPoolConnections),
		zap.Int("maxIdleConns", maxIdleConns))
	db.SetMaxIdleConns(maxIdleConns)
	return db, err
}

// initFrabitDB attempts to create/upgrade the frabit backend database. It is created once in the
// application's lifetime.
func initFrabitDB(db *sql.DB) error {
	deployStatements(db, generateSQLBase)
	deployStatements(db, generateSQLPatches)

	return nil
}

// deployStatements will issue given sql queries that are not already known to be deployed.
// This iterates both lists (to-run and already-deployed) and also verifies no contraditions.
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
