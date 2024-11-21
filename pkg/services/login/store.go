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

package login

import (
	"context"
	"github.com/frabits/frabit/pkg/common/utils"
	"gorm.io/gorm"
	"time"
)

type Store interface {
	AddRecord(ctx context.Context, cmd *LoginAttempt) error
	DeleteLoginAttempt(ctx context.Context, login string) error
	DeleteOlderThanLoginAttempt(ctx context.Context, olderThan time.Time) error
	GetUserLoginAttemptCount(ctx context.Context, login string) (int64, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func providerStore(db *gorm.DB) Store {
	return &storeImpl{DB: db}
}

func (s *storeImpl) AddRecord(ctx context.Context, cmd *LoginAttempt) error {
	result := s.DB.Create(&LoginAttempt{
		Login:     cmd.Login,
		ClientIP:  cmd.ClientIP,
		CreatedAt: cmd.CreatedAt,
	})
	return result.Error
}

func (s *storeImpl) DeleteLoginAttempt(ctx context.Context, login string) error {
	s.DB.Model(&LoginAttempt{}).Where("login=?", login).Update("last_seen_at", utils.NowDatetime())
	return nil
}

func (s *storeImpl) DeleteOlderThanLoginAttempt(ctx context.Context, olderThan time.Time) error {
	s.DB.Where("create_at<?", olderThan).Delete(&LoginAttempt{})
	return nil
}

func (s *storeImpl) GetUserLoginAttemptCount(ctx context.Context, login string) (int64, error) {
	var loginCount int64
	s.DB.Model(&LoginAttempt{}).Where("login=?", login).Count(&loginCount)
	return loginCount, nil
}
