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
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"
	"time"
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

type loginAttemptImpl struct {
	cfg   *config.Config
	log   *slog.Logger
	store Store
}

func ProviderLoginNative(conf *config.Config, metaDB *db.MetaStore) Service {
	store := providerStore(metaDB.Gdb)
	ls := &loginAttemptImpl{
		cfg:   conf,
		log:   log.New("login"),
		store: store,
	}
	return ls
}

func (s *loginAttemptImpl) Run(ctx context.Context) error {
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			tick.Stop()
			return ctx.Err()
		case <-tick.C:
			s.cleanup(ctx)
		}
	}
}

func (s *loginAttemptImpl) Add(ctx context.Context, cmd *CreateLoginAttemptCmd) error {
	now := time.Now()
	la := &LoginAttempt{
		Login:     cmd.Login,
		ClientIP:  cmd.ClientIP,
		CreatedAt: now.Format(time.DateTime),
	}
	return s.store.AddRecord(ctx, la)
}

func (s *loginAttemptImpl) Reset(ctx context.Context, login string) error {
	return s.store.DeleteLoginAttempt(ctx, login)
}

func (s *loginAttemptImpl) Validate(ctx context.Context, login string) bool {
	if s.cfg.Server.DisableLoginProtection {
		return true
	}
	loginCount, err := s.store.GetUserLoginAttemptCount(ctx, login)
	if err != nil {
		return false
	}
	if loginCount > s.cfg.Server.LoginMaxRetry {
		return false
	}
	return true
}

func (s *loginAttemptImpl) cleanup(ctx context.Context) {
	olderThan := time.Now().Add(-10 * time.Minute)
	s.store.DeleteOlderThanLoginAttempt(ctx, olderThan)
}
