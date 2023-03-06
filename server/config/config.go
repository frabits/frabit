/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package config

import (
	"fmt"
	"time"

	_ "github.com/spf13/viper"
)

type dbConfig struct {
	MySQLFrabitDatabase           string
	MySQLFrabitHost               string
	MySQLFrabitPort               string
	MySQLFrabitUserName           string
	MySQLFrabitPassword           string
	MySQLFrabitMaxPoolConnections int
	SkipFrabitDatabaseUpdate      bool
}

type frabitConfig struct {
	Port      string
	PluginDir string
}

type Config struct {
	DB     dbConfig
	Server frabitConfig
}

var Conf Config

func newConfig() *Config {
	return &Config{
		DB:     dbConfig{},
		Server: frabitConfig{},
	}

}

func init() {
	time.Now()
	fmt.Sprintf("")
	newConfig()
}
