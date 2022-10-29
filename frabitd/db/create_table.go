/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package db

// generateSQLBase & generateSQLPatches are lists of SQL statements required to build the orchestrator backend
var generateSQLBase = []string{
	`
        CREATE TABLE IF NOT EXISTS instance (
          hostname varchar(128) CHARACTER SET ascii NOT NULL,
          port smallint(5) unsigned NOT NULL,
          last_checked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          last_seen timestamp NULL DEFAULT NULL,
          server_id int(10) unsigned NOT NULL,
          version varchar(128) CHARACTER SET ascii NOT NULL,
          binlog_format varchar(16) CHARACTER SET ascii NOT NULL,
          log_bin tinyint(3) unsigned NOT NULL,
          log_slave_updates tinyint(3) unsigned NOT NULL,
          binary_log_file varchar(128) CHARACTER SET ascii NOT NULL,
          binary_log_pos bigint(20) unsigned NOT NULL,
          PRIMARY KEY (hostname,port)
        ) ENGINE=InnoDB	`,
	`
		CREATE INDEX first_seen_idx_database_instance_stale_binlog_coordinates ON database_instance_stale_binlog_coordinates (first_seen)
	`,
}
