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

package user

import (
	"context"
	"github.com/frabits/frabit/pkg/common/utils"
	"gorm.io/gorm"
)

type Store interface {
	CreateUser(ctx context.Context, user *User) (uint32, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByLogin(context.Context, string) (User, error)
	GetUserById(context.Context, uint32) (User, error)
	GetUserServiceAccount(context.Context) ([]User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, user *User) error
	UpdateUserLastSeen(ctx context.Context, login string) error
	VerifyEmail(ctx context.Context, login string) error
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{db}
}

func (s *storeImpl) CreateUser(ctx context.Context, user *User) (uint32, error) {
	s.DB.Create(user)
	return 0, nil
}

func (s *storeImpl) DeleteUser(ctx context.Context, login string) error {
	var user User
	s.DB.Where("login=?", login).Delete(user)
	return nil
}

func (s *storeImpl) UpdateUser(ctx context.Context, user *User) error {
	s.DB.Model(&User{}).Where("login=?", user.Login).Updates(User{})
	return nil
}

func (s *storeImpl) VerifyEmail(ctx context.Context, login string) error {
	s.DB.Model(&User{}).Where("login=?", login).Update("is_email_verified", 1)
	return nil
}

func (s *storeImpl) UpdateUserLastSeen(ctx context.Context, login string) error {
	s.DB.Model(&User{}).Where("login=?", login).Update("last_seen_at", utils.NowDatetime())
	return nil
}

func (s *storeImpl) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	s.DB.Model(User{}).Find(&users)
	return users, nil
}

func (s *storeImpl) GetUserByLogin(ctx context.Context, login string) (User, error) {
	var user User
	s.DB.Model(User{}).Where("login=?", login).First(&user)
	return user, nil
}

func (s *storeImpl) GetUserById(ctx context.Context, uid uint32) (User, error) {
	var user User
	s.DB.Model(User{}).Where("id=?", uid).First(&user)
	return user, nil
}

func (s *storeImpl) GetUserServiceAccount(context.Context) ([]User, error) {
	var users []User
	s.DB.Model(User{}).Where("is_service_account=?", 1).Find(&users)
	return users, nil
}
