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
	"encoding/hex"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
	"unicode"

	fb "github.com/frabits/frabit-go-sdk/frabit"

	"github.com/frabits/frabit/pkg/common/constant"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/secrets"
)

type service struct {
	log    *slog.Logger
	cfg    *config.Config
	secret secrets.Service
	store  Store
}

func ProviderService(conf *config.Config, metaDB *db.MetaStore, secret secrets.Service) Service {
	store := NewStoreImpl(metaDB.Gdb)
	us := &service{
		log:    log.New("user"),
		cfg:    conf,
		secret: secret,
		store:  store,
	}

	return us
}

func (s *service) CreateUser(ctx context.Context, createReq fb.CreateUserRequest) (uint32, error) {
	if err := s.ValidatePassword(ctx, createReq.Password); err != nil {
		s.log.Error("cannot create user", "Error", err.Error())
		return 0, err
	}

	user := &User{
		Login:            createReq.Login,
		Username:         createReq.Name,
		Email:            createReq.Email,
		Password:         utils.GeneratePassword(createReq.Password),
		Rands:            utils.GenRandom(12),
		IsAdmin:          0,
		IsDisabled:       1,
		IsExternal:       0,
		IsEmailVerified:  0,
		IsServiceAccount: createReq.IsServiceAccount,
		Theme:            createReq.Theme,
		OrgId:            constant.GlobalOrgId,
		CreatedAt:        utils.NowDatetime(),
		UpdatedAt:        utils.NowDatetime(),
		LastSeenAt:       utils.NowDatetime(),
	}
	// service account can be empty password
	if user.IsServiceAccount == 1 {
		user.Password = ""
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

func (s *service) UpdateUser(ctx context.Context, req fb.UpdateUserRequest) error {
	existUser, err := s.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return err
	}
	if err := s.ValidatePassword(ctx, req.Password); err != nil {
		s.log.Error("cannot update user", "Error", err.Error())
		return err
	}
	existUser.Password = utils.GeneratePassword(req.Password)

	return s.store.UpdateUser(ctx, existUser)
}

func (s *service) GetServiceAccount(ctx context.Context) ([]UserProfileDTO, error) {
	var users []UserProfileDTO
	saAccounts, err := s.store.GetUserServiceAccount(ctx)
	if err != nil {
		s.log.Error("query user failed", "Error", err.Error())
	}
	for _, account := range saAccounts {
		user := UserProfileDTO{
			Login:           account.Login,
			Username:        account.Username,
			Email:           account.Email,
			Password:        account.Password,
			IsAdmin:         account.IsAdmin,
			IsDisabled:      account.IsDisabled,
			IsExternal:      account.IsExternal,
			IsEmailVerified: account.IsEmailVerified,
			Theme:           account.Theme,
			OrgId:           account.OrgId,
			CreatedAt:       account.CreatedAt,
			UpdatedAt:       account.UpdatedAt,
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *service) GetUserByLogin(ctx context.Context, login string) (*User, error) {
	userInfo, err := s.store.GetUserByLogin(ctx, login)
	if err != nil {
		s.log.Error("query user failed", "Error", err.Error())
	}
	return &userInfo, nil
}

func (s *service) GetUserById(ctx context.Context, uid uint32) (UserProfileDTO, error) {
	var user UserProfileDTO
	users, err := s.store.GetUserById(ctx, uid)
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

func (s *service) GenerateCode(ctx context.Context, login string) (string, error) {
	user, err := s.GetUserByLogin(ctx, login)
	if err != nil {
		s.log.Error("get user detail failed", "Error", err.Error())
		return "", err
	}
	// generate verify code based user
	validDatetime := time.Now().Add(5 * time.Minute).Unix()
	code := fmt.Sprintf("%s:%s:%s:%v", user.Login, user.Email, user.Rands, validDatetime)
	encCode, _ := s.secret.Encrypt([]byte(code))
	encodedCode := hex.EncodeToString(encCode)
	return encodedCode, nil
}

func (s *service) ValidateCode(ctx context.Context, code string) ([]string, error) {
	decodedCode, _ := hex.DecodeString(code)
	decCode, _ := s.secret.Decrypt(decodedCode)
	parts := strings.Split(string(decCode), ":")
	if len(parts) != 4 {
		return nil, ErrVerifyCodeInvalid
	}

	// check expire
	validDatetime, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		return nil, ErrVerifyCodeExpired
	}
	now := time.Now().Unix()
	if now > validDatetime {
		return nil, ErrVerifyCodeExpired
	}
	return parts, nil
}

func (s *service) ValidatePassword(ctx context.Context, password string) error {
	if !s.cfg.Server.BasicAuthStrongPolicy {
		if len(password) < 4 {
			return ErrPasswordTooShort
		}
		return nil
	}

	if len(password) < minPasswordLength {
		return ErrPasswordTooShortForStrongPolicy
	}

	hasNumber := false
	hasUpperCase := false
	hasLowerCase := false
	hasSymbol := false
	for _, r := range password {
		if !hasNumber && unicode.IsNumber(r) {
			hasNumber = true
		}

		if !hasUpperCase && unicode.IsUpper(r) {
			hasUpperCase = true
		}

		if !hasLowerCase && unicode.IsLower(r) {
			hasLowerCase = true
		}

		if !hasNumber && !hasUpperCase && !hasLowerCase {
			hasSymbol = true
		}

		if hasNumber && hasUpperCase && hasLowerCase && hasSymbol {
			return nil
		}
	}

	return ErrPasswordNotMatchStrongPolicy
}

func (s *service) DeleteUser(ctx context.Context, login string) error {
	return s.store.DeleteUser(ctx, login)
}

func (s *service) UpdateUserLastSeen(ctx context.Context, login string) error {
	s.log.Info("")
	err := s.store.UpdateUserLastSeen(ctx, login)
	if err != nil {
		s.log.Error("update user last seen failed", "Error", err.Error())
		return err
	}
	s.log.Debug("update user last seen successfully")
	return nil
}

func (s *service) VerifyEmail(ctx context.Context, login string) error {
	s.log.Info("start verify email")
	err := s.store.VerifyEmail(ctx, login)
	if err != nil {
		s.log.Error("update user last seen failed", "Error", err.Error())
		return err
	}
	s.log.Debug("update user last seen successfully")
	return nil
}
