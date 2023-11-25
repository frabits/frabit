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

package config

import (
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
	newConfig()
}
