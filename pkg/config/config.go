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
	"log/slog"
)

type dbConfig struct {
	Database           string
	Host               string
	Port               uint32
	UserName           string
	Password           string
	MaxPoolConnections int
	DatabaseUpdate     bool
}

type frabitConfig struct {
	Port         uint32
	PluginDir    string
	FileName     string
	Format       string
	DefaultLevel slog.Level
	MaxDay       uint
}

type Config struct {
	DB     dbConfig
	Server frabitConfig
}

var Conf = newConfig()

func newConfig() *Config {
	dbConf := dbConfig{
		Database:           "frabit",
		Host:               "127.0.0.1",
		Port:               3306,
		UserName:           "frabit",
		Password:           "frabitSecurePasswd",
		MaxPoolConnections: 20,
	}
	frabitConf := frabitConfig{
		Port:         9180,
		FileName:     "/tmp/frabit.log",
		Format:       "json",
		DefaultLevel: slog.LevelInfo,
	}
	return &Config{
		DB:     dbConf,
		Server: frabitConf,
	}
}
