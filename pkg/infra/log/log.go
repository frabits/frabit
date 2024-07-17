// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
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

package log

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/frabits/frabit/pkg/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

var globalLogger *slog.Logger

type FileLogger struct {
	fileName     string
	format       string
	maxDay       int64
	defaultLevel slog.Level
	log          *slog.Logger
	fd           *os.File
}

func New(name string) *slog.Logger {
	return globalLogger.With("logger", name)
}

func NewLogger(conf *config.Config) *FileLogger {
	logfile := conf.Server.FileName
	logFormat := conf.Server.Format
	if logfile == "" {
		logfile = "/tmp/frabit.log"
	}
	if err := os.MkdirAll(filepath.Dir(logfile), 0744); err != nil {
		fmt.Printf("log dir create filed:%s\n", err.Error())
	}

	if strings.ToLower(logFormat) == "" {
		logFormat = "json"
	}

	logger := &FileLogger{
		fileName:     logfile,
		format:       logFormat,
		defaultLevel: conf.Server.DefaultLevel,
		maxDay:       3,
	}

	return logger
}

func (fl *FileLogger) init() {
	fd, err := fl.createLogfile()
	if err != nil {
		fmt.Printf("create logfile failed, err:%s\n", err.Error())
	}
	fl.fd = fd
}

func (fl *FileLogger) createLogfile() (*os.File, error) {
	return os.OpenFile(fl.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
}

func (fl *FileLogger) StartLogger() {
	fl.init()
	dest := &lumberjack.Logger{
		Filename:   fl.fileName,
		MaxSize:    50,
		MaxBackups: 5,
		Compress:   true,
	}
	switch strings.ToLower(fl.format) {
	case "json":
		fl.log = slog.New(slog.NewJSONHandler(dest, &slog.HandlerOptions{AddSource: fl.defaultLevel == slog.LevelDebug, Level: fl.defaultLevel}))
	default:
		fl.log = slog.New(slog.NewTextHandler(dest, &slog.HandlerOptions{AddSource: fl.defaultLevel == slog.LevelDebug, Level: fl.defaultLevel}))
	}

	globalLogger = fl.log
}

func init() {
	logger := NewLogger(config.Conf)
	logger.StartLogger()
}
