/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package config

type Config struct {
	MySQLFrabitDatabase           string
	MySQLFrabitHost               string
	MySQLFrabitPort               string
	MySQLFrabitUserName           string
	MySQLFrabitPassword           string
	MySQLFrabitMaxPoolConnections int
	SkipFrabitDatabaseUpdate      bool
}

var Conf Config
