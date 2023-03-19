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

package service

import (
	"context"
	"time"

	"github.com/frabits/frabit/common/log"
	pxb "github.com/frabits/frabit/pkg/xtrabackup"

	"go.uber.org/zap"
)

// MySQLBackupService implement DB backup task
type MySQLBackupService struct {
	StartDatetime time.Time
	EndDatetime   time.Time
	PXB           pxb.Xtrabackup
	Type          BackupType
	Logger        *zap.Logger
}

func newBackupService() *MySQLBackupService {
	log.Info("create BackupService")
	return &MySQLBackupService{
		Logger: log.Logger,
	}
}

// Start a backup job
func (bak *MySQLBackupService) Start(ctx context.Context, bakType BackupType) error {
	bak.Logger.Info("create BackupService")
	bak.Type = bakType
	bak.PXB.Backup()
	return nil
}

func (bak *MySQLBackupService) Stop() error {

	return nil
}

func (bak *MySQLBackupService) Cancel() error {

	return nil
}

func init() {
	newBackupService()
}
