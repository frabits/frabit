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
	"github.com/frabits/frabit/pkg/services/login"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/apikey"
	"github.com/frabits/frabit/pkg/services/secrets"
	"github.com/frabits/frabit/pkg/services/user"
)

type AuthnImpl struct {
	log     *slog.Logger
	cfg     *config.Config
	apikey  apikey.Service
	user    user.Service
	secrets secrets.Service
	clients map[string]Client
}

func ProviderService(cfg *config.Config, loginAttempt login.Service, apikey apikey.Service, user user.Service, secrets secrets.Service) Service {
	ai := &AuthnImpl{
		log:     log.New("authn"),
		cfg:     cfg,
		apikey:  apikey,
		user:    user,
		secrets: secrets,
		clients: make(map[string]Client, 0),
	}
	// register all client based config
	ai.RegisterClient(ProviderApiKey(apikey, user))
	var passwdClient []PasswordClient
	if !cfg.Server.DisableBasic {
		frabit := ProviderFrabit(user)
		passwdClient = append(passwdClient, frabit)
	}
	if !cfg.Server.DisableLdap {
		ldap := ProviderLdap()
		passwdClient = append(passwdClient, ldap)
	}

	if len(passwdClient) > 0 {
		pwClient := ProviderPassword(loginAttempt, passwdClient...)
		basic := ProviderBasic(pwClient)
		form := ProviderForm(pwClient)
		ai.RegisterClient(basic)
		ai.RegisterClient(form)
	}
	ai.log.Info("register all session client", "Total number", len(ai.clients))
	// all oauth client dynamic enable via web console
	return ai
}

func (s *AuthnImpl) Login(ctx context.Context, client string, authReq *AuthRequest) (*Identity, error) {
	authClient, ok := s.clients[client]
	if !ok {
		return nil, nil
	}
	s.log.Info("found auth client", "Name", authClient.Name())
	return authClient.Authenticate(ctx, authReq)
}

func (s *AuthnImpl) Logout(ctx context.Context, session string) (*Redirect, error) {
	redirect := &Redirect{
		URL: "/login",
	}
	// remove session

	return redirect, nil
}

func (s *AuthnImpl) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 60)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			s.log.Info("Start purge dynamic disabled oauth client")
		}
	}
}

func (s *AuthnImpl) RegisterClient(client Client) {
	s.clients[client.Name()] = client
}

func (s *AuthnImpl) UnRegisterClient(client Client) {
	delete(s.clients, client.Name())
}
