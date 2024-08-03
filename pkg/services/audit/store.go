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

package audit

import (
	"context"

	"gorm.io/gorm"
)

type Store interface {
	Create(ctx context.Context, event *AuditLog) error
	GetEvent(ctx context.Context) ([]AuditLog, error)
	GetEventByUsername(ctx context.Context, username string) ([]*AuditLog, error)
	GetEventByEvent(ctx context.Context, eventName string) ([]*AuditLog, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) Store {
	return &storeImpl{DB: db}
}

func (s *storeImpl) Create(ctx context.Context, event *AuditLog) error {
	s.DB.Create(event)
	return nil
}

func (s *storeImpl) GetEvent(ctx context.Context) ([]AuditLog, error) {
	var auditLogs []AuditLog
	s.DB.Model(AuditLog{}).Find(&auditLogs)
	return auditLogs, nil
}

func (s *storeImpl) GetEventByUsername(ctx context.Context, username string) ([]*AuditLog, error) {
	var auditLogs []*AuditLog
	return auditLogs, nil
}

func (s *storeImpl) GetEventByEvent(ctx context.Context, username string) ([]*AuditLog, error) {
	var auditLogs []*AuditLog
	return auditLogs, nil
}
