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

// generateSQLBase & generateSQLPatches are lists of SQL statements required to build the frabit backend
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
        CREATE TABLE IF NOT EXISTS cluster (
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
        CREATE TABLE IF NOT EXISTS license (
          id bigint NOT NULL auto_increment,
          license_level varchar(15) NOT NULL default "basic" comment "valid license level include：community、enterprise and ultimate ",
          current_license varchar(500) NOT NULL DEFAULT '',
          last_license varchar(500) NOT NULL DEFAULT '',
          create_at varchar(50) not null default "" comment '',
          update_at varchar(50) not null default "" comment '',
          PRIMARY KEY (id)
        ) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS version (
	      id bigint NOT NULL auto_increment,
	      version varchar(15) NOT NULL default "v1.0.0" comment "frabit component version", 
          create_at varchar(50) not null default "",
          update_at varchar(50) not null default "",
	      PRIMARY KEY (id)
	) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS user (
	      id bigint NOT NULL auto_increment,
	      username varchar(100) NOT NULL default "" comment "login user name", 
          email varchar(100) NOT NULL default "", 
          password varchar(200) not null ,
          is_disabled tinyint not null default 0,
	      PRIMARY KEY (id)
	) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS team (
	      id bigint NOT NULL auto_increment,
	      name varchar(100) NOT NULL default "" comment "login user name", 
          description varchar(200) NOT NULL default "", 
          password varchar(200) not null ,
          is_disabled tinyint not null default 0,
	      PRIMARY KEY (id)
	) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS team_member (
	      id bigint NOT NULL auto_increment,
	      team_id varchar(100) NOT NULL default "" comment "login user name", 
          user_id varchar(200) NOT NULL default "",
	      PRIMARY KEY (id),
          UNIQUE KEY uniq_tu (team_id,user_id)
	) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS project (
	      id bigint NOT NULL auto_increment,
	      user_name varchar(100) NOT NULL default "" comment "login user name", 
          email varchar(100) NOT NULL default "", 
	      PRIMARY KEY (id)
	) ENGINE=InnoDB`,
	`
        CREATE TABLE IF NOT EXISTS workspace (
	      id bigint NOT NULL auto_increment,
	      user_name varchar(100) NOT NULL default "" comment "login user name", 
          email varchar(100) NOT NULL default "", 
	      PRIMARY KEY (id)
	) ENGINE=InnoDB`,
}
