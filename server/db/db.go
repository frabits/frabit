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
)

// OpenTopology returns the DB instance for the frabit backed database
func OpenOrchestrator() (db *sql.DB, err error) {
	// first time ever we talk to MySQL
	query := fmt.Sprintf("create database if not exists %s", config.Conf.MySQLFrabitDatabase)
	if _, err := db.Exec(query); err != nil {
		return db, log.Errore(err)
	}
	if !config.Config.SkipOrchestratorDatabaseUpdate {
		initFrabitDB(db)
	}
	maxIdleConns := int(100)
	if maxIdleConns < 10 {
		maxIdleConns = 10
	}
	log.Infof("Connecting to backend %s:%d: maxConnections: %d, maxIdleConns: %d",
		config.Conf.MySQLFrabitHost,
		config.Conf.MySQLFrabitPort,
		config.Conf.MySQLFrabitMaxPoolConnections,
		maxIdleConns)
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
		log.Fatale(err)
	}
	for i, query := range queries {
		if i == 0 {
			//log.Debugf("sql_mode is: %+v", originalSqlMode)
		}
		if _, err := tx.Exec(query); err != nil {
			if strings.Contains(err.Error(), "syntax error") {
				return log.Fatalf("Cannot initiate orchestrator: %+v; query=%+v", err, query)
			}
			if !strings.Contains(err.Error(), "duplicate column name") &&
				!strings.Contains(err.Error(), "Duplicate column name") &&
				!strings.Contains(err.Error(), "check that column/key exists") &&
				!strings.Contains(err.Error(), "already exists") &&
				!strings.Contains(err.Error(), "Duplicate key name") {
				log.Errorf("Error initiating frabit: %+v; query=%+v", err, query)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatale(err)
	}
	return nil
}
