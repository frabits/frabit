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
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/notifications"
	"github.com/frabits/frabit/pkg/services/secrets"
	"log/slog"
)

type Verifier interface {
	Start(ctx context.Context, username string) error
	Complete(ctx context.Context, code string) error
}

type verifyImpl struct {
	cfg         *config.Config
	log         *slog.Logger
	ns          *notifications.Service
	userService Service
	secret      secrets.Service
}

func ProviderVerifier(cfg *config.Config, ns *notifications.Service, userService Service, secret secrets.Service) Verifier {
	vi := &verifyImpl{
		cfg:         cfg,
		log:         log.New("verify"),
		ns:          ns,
		userService: userService,
		secret:      secret,
	}

	return vi
}

func (s *verifyImpl) Start(ctx context.Context, username string) error {
	s.log.Info("start verify email", "username", username)
	code, err := s.userService.GenerateCode(ctx, username)
	if err != nil {
		s.log.Error("cannot generate code", "Error", err.Error())
		return err
	}
	s.log.Info("generate verify code", "code", code)
	return nil
}

func (s *verifyImpl) Complete(ctx context.Context, code string) error {
	parts, err := s.userService.ValidateCode(ctx, code)
	if err != nil {
		s.log.Error("validate code failed", "Error", err.Error())
		return err
	}
	return s.userService.VerifyEmail(ctx, parts[0])
}
