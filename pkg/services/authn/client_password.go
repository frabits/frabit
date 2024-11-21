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

package authn

import (
	"context"
	"errors"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"
	"net/http"

	"github.com/frabits/frabit/pkg/services/login"
)

type Password struct {
	log          *slog.Logger
	loginAttempt login.Service
	clients      []PasswordClient
}

func ProviderPassword(loginAttempt login.Service, clients ...PasswordClient) PasswordClient {
	return &Password{
		log:          log.New("authn.client.password"),
		loginAttempt: loginAttempt,
		clients:      clients,
	}
}

func (c *Password) AuthenticatePasswd(ctx context.Context, req *http.Request, username string, password string) (*Identity, error) {
	var authErrs error
	allowedLogin := c.loginAttempt.Validate(ctx, username)
	if !allowedLogin {
		return nil, ErrFailedLoginTooMuch
	}
	for _, pdClient := range c.clients {
		id, authErr := pdClient.AuthenticatePasswd(ctx, req, username, password)
		if authErr != nil {
			if errors.Is(authErr, ErrCredential) {
				la := &login.CreateLoginAttemptCmd{
					Login:    username,
					ClientIP: req.RemoteAddr,
				}
				if err := c.loginAttempt.Add(ctx, la); err != nil {
					c.log.Warn("cannot add login attempt")
				}
			}
			authErrs = errors.Join(authErr)
			continue
		}
		return id, nil
	}

	if errors.Is(authErrs, nil) {

	}

	return nil, nil
}
