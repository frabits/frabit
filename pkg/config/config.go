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
	"log/slog"
	"time"
)

var (
	Security string
	Logger   LoggerConfig
)

type dbConfig struct {
	Database           string
	Host               string
	Port               uint32
	UserName           string
	Password           string
	MaxPoolConnections int
	SkipDatabaseUpdate bool
}

type LoggerConfig struct {
	FileName     string
	Format       string
	Security     string
	DefaultLevel slog.Level
	MaxDay       uint
}

type frabitConfig struct {
	Port                     uint32
	PluginDir                string
	SecureKey                string
	LoginMaxInactiveLifetime time.Duration
	LoginMaxLifetime         time.Duration
}

type Config struct {
	DB     dbConfig
	Server frabitConfig
	Logger LoggerConfig
}

func ProviderConfig() *Config {
	dbConf := dbConfig{
		Database:           "frabit",
		Host:               "192.168.1.7",
		Port:               3306,
		UserName:           "frabit",
		Password:           "frabitSecure_1Passwd",
		MaxPoolConnections: 20,
		SkipDatabaseUpdate: false,
	}

	Logger = LoggerConfig{
		FileName:     "/tmp/frabit.log",
		Format:       "json",
		DefaultLevel: slog.LevelInfo,
	}
	Security = "2xWkMhUhF96IVtawQ7mRHuU6uYB"
	frabitConf := frabitConfig{
		Port:      9180,
		SecureKey: Security,
	}

	return &Config{
		DB:     dbConf,
		Server: frabitConf,
		Logger: Logger,
	}
}
