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
	"log/slog"

	fb "github.com/frabits/frabit-go-sdk/frabit"
	"github.com/frabits/frabit/pkg/common/constant"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
)

type service struct {
	log   *slog.Logger
	cfg   *config.Config
	store Store
}

func ProviderService(conf *config.Config, metaDB *db.MetaStore) Service {
	store := NewStoreImpl(metaDB.Gdb)
	us := &service{
		log:   log.New("user"),
		cfg:   conf,
		store: store,
	}

	return us
}

func (s *service) CreateUser(ctx context.Context, createReq fb.CreateUserRequest) (uint32, error) {
	user := &User{
		Login:           createReq.Login,
		Username:        createReq.Name,
		Email:           createReq.Email,
		Password:        utils.GeneratePassword(createReq.Password),
		Rands:           utils.GenRandom(12),
		IsAdmin:         0,
		IsDisabled:      1,
		IsExternal:      0,
		IsEmailVerified: 0,
		Theme:           createReq.Theme,
		OrgId:           constant.GlobalOrgId,
		CreatedAt:       utils.NowDatetime(),
		UpdatedAt:       utils.NowDatetime(),
		LastSeenAt:      utils.NowDatetime(),
	}
	uid, err := s.store.CreateUser(ctx, user)
	if err != nil {
		s.log.Error("create user failed", "Error", err.Error())
		return 0, err
	}
	return uid, nil
}

func (s *service) GetUsers(ctx context.Context) ([]User, error) {
	return s.store.GetUsers(ctx)
}

func (s *service) UpdateUser(ctx context.Context) error {
	return nil
}

func (s *service) GetUserByLogin(ctx context.Context, login string) (UserProfileDTO, error) {
	var user UserProfileDTO
	users, err := s.store.GetUserByLogin(ctx, login)
	if err != nil {
		s.log.Error("query user failed", "Error", err.Error())
	}
	user = UserProfileDTO{
		Login:           users.Login,
		Username:        users.Username,
		Email:           users.Email,
		Password:        users.Password,
		IsAdmin:         user.IsAdmin,
		IsDisabled:      users.IsDisabled,
		IsExternal:      user.IsExternal,
		IsEmailVerified: users.IsEmailVerified,
		Theme:           users.Theme,
		OrgId:           users.OrgId,
		CreatedAt:       users.CreatedAt,
		UpdatedAt:       users.UpdatedAt,
	}
	return user, nil
}

func (s *service) DeleteUser(ctx context.Context, login string) error {
	return s.store.DeleteUser(ctx, login)
}

func (s *service) UpdateUserLastSeen(ctx context.Context, login string) error {
	s.log.Info("")
	err := s.store.UpdateUser(ctx, login)
	if err != nil {
		s.log.Error("update user last seen failed", "Error", err.Error())
		return err
	}
	s.log.Debug("update user last seen successfully")
	return nil
}
