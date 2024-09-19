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

package auth

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type store interface {
	CreateToken(ctx context.Context, auth *UserAuth) error
	DeleteTokenByTime(ctx context.Context, createdBefore time.Time, rotatedBefore time.Time) error
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{DB: db}
}

func (s *storeImpl) CreateToken(ctx context.Context, auth *UserAuth) error {
	s.DB.Create(auth)
	return nil
}

func (s *storeImpl) DeleteTokenByTime(ctx context.Context, createdBefore time.Time, rotatedBefore time.Time) error {
	var authToken []UserAuth
	s.DB.Where("created_at <= ? and rotated_at <=?", createdBefore.Format(time.DateTime), rotatedBefore.Format(time.DateTime)).Delete(&authToken)
	return nil
}
