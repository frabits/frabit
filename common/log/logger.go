// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
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
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	gLevel zap.AtomicLevel
)

// Initializes the global console logger.
func init() {
	gLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	Logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		gLevel,
	))
}

// SetLevel wraps the zap Level's SetLevel method.
func SetLevel(level zapcore.Level) {
	gLevel.SetLevel(level)
}

// EnabledLevel wraps the zap Level's Enabled method.
func EnabledLevel(level zapcore.Level) bool {
	return gLevel.Enabled(level)
}

// Debug wraps the zap Logger's Debug method.
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info wraps the zap Logger's Info method.
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn wraps the zap Logger's Warn method.
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error wraps the zap Logger's Error method.
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Sync wraps the zap Logger's Sync method.
func Sync() {
	_ = Logger.Sync()
}
