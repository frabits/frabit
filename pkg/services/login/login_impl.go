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
	"errors"
	"log/slog"
	"strings"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/user"
)

var (
	ErrEmailNotAllowed       = errors.New("required email domain not fulfilled")
	ErrInvalidCredentials    = errors.New("invalid username or password")
	ErrNoEmail               = errors.New("login provider didn't return an email address")
	ErrProviderDeniedRequest = errors.New("login provider denied login request")
	ErrTooManyLoginAttempts  = errors.New("too many consecutive incorrect login attempts for user - login for user temporarily blocked")
	ErrPasswordEmpty         = errors.New("no password provided")
	ErrUserDisabled          = errors.New("user is disabled")
	ErrAbsoluteRedirectTo    = errors.New("absolute URLs are not allowed for redirect_to cookie value")
	ErrInvalidRedirectTo     = errors.New("invalid redirect_to cookie value")
	ErrForbiddenRedirectTo   = errors.New("forbidden redirect_to cookie value")
	ErrNoAuthProvider        = errors.New("enable at least one login provider")
)

type LoginNative struct {
	cfg         *config.Config
	log         *slog.Logger
	userService user.Service
	store       Store
}

func ProviderLoginNative(conf *config.Config, user user.Service) Service {
	ls := &LoginNative{
		cfg:         conf,
		log:         log.New("login"),
		userService: user,
	}
	return ls
}

func (s *LoginNative) Authenticator(ctx context.Context, authInfo AuthPasswd) error {
	s.log.Debug("Auth user via username and password")
	login := strings.ToLower(authInfo.Login)
	userInfo, err := s.userService.GetUserByLogin(ctx, login)
	if err != nil {
		return err
	}
	if ok := utils.ComparePassword(authInfo.Password, userInfo.Password); !ok {
		s.log.Error(ErrInvalidCredentials.Error())
		return ErrInvalidCredentials
	}
	if err := s.userService.UpdateUserLastSeen(ctx, login); err != nil {
		s.log.Error(ErrInvalidCredentials.Error())
	}
	return nil
}
