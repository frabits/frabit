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

package session

import (
	"context"
	"github.com/frabits/frabit/pkg/common/utils"
	"time"

	"gorm.io/gorm"
)

type store interface {
	CreateSession(ctx context.Context, auth *Session) error
	GetSessionByToken(ctx context.Context, token string) (*Session, error)
	RevokeUserAllSessions(ctx context.Context, login string) error
	DeleteSessionByTime(ctx context.Context, createdBefore time.Time, rotatedBefore time.Time) error
	DeleteSessionById(ctx context.Context, id uint32) error
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{DB: db}
}

func (s *storeImpl) CreateSession(ctx context.Context, session *Session) error {
	s.DB.Create(&Session{
		Login:     session.Login,
		ClientIP:  session.ClientIP,
		UserAgent: session.UserAgent,
		Token:     session.Token,
		CreatedAt: session.CreatedAt,
		UpdatedAt: session.UpdatedAt,
		RotatedAt: session.RotatedAt,
	})
	return nil
}

func (s *storeImpl) GetSessionByToken(ctx context.Context, token string) (*Session, error) {
	var session Session
	s.DB.Where("token=?", token).Find(&session)
	return &session, nil
}

func (s *storeImpl) RevokeUserAllSessions(ctx context.Context, login string) error {
	s.DB.Model(&Session{}).Where("login=?", login).Update("last_seen_at", utils.NowDatetime())
	return nil
}

func (s *storeImpl) DeleteSessionByTime(ctx context.Context, createdBefore time.Time, rotatedBefore time.Time) error {
	var authToken []Session
	s.DB.Where("created_at <= ? and rotated_at <=?", createdBefore.Format(time.DateTime), rotatedBefore.Format(time.DateTime)).Delete(&authToken)
	return nil
}

func (s *storeImpl) DeleteSessionById(ctx context.Context, id uint32) error {
	var authToken []Session
	s.DB.Where("id=?", id).Delete(&authToken)
	return nil
}
