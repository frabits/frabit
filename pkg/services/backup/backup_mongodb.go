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

package backup

import (
	"context"
	"github.com/frabits/frabit/pkg/common/log"
	"time"

	pxb "github.com/frabits/frabit/pkg/xtrabackup"

	"go.uber.org/zap"
)

// MongoBackup implement DB backup task
type MongoBackup struct {
	StartDatetime time.Time
	EndDatetime   time.Time
	PXB           pxb.Xtrabackup
	Type          BackupType
	Logger        *zap.Logger
}

func newMongoBackup() *MongoBackup {
	// log.Info("create BackupService")
	return &MongoBackup{
		Logger: log.Logger,
	}
}

// CreateBackup a backup job
func (bak *MongoBackup) CreateBackup(ctx context.Context) error {
	bak.Logger.Info("create BackupService")
	bak.PXB.Backup()
	return nil
}

// ListBackup a backup job
func (bak *MongoBackup) ListBackup(ctx context.Context) ([]BackupSet, error) {
	bak.Logger.Info("create BackupService")
	bak.PXB.Backup()
	bs := make([]BackupSet, 0)
	return bs, nil
}

// CancelBackup a backup job
func (bak *MongoBackup) CancelBackup(ctx context.Context) error {
	bak.Logger.Info("create BackupService")
	bak.PXB.Backup()
	return nil
}

// PurgeBackup a backup job
func (bak *MongoBackup) PurgeBackup(ctx context.Context) error {
	bak.Logger.Info("create BackupService")
	bak.PXB.Backup()
	return nil
}

func init() {
	newMongoBackup()
}
