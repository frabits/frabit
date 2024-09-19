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
	"log/slog"

	"github.com/frabits/frabit/pkg/infra/log"
)

var (
	ErrBasic = errors.New("")
)

type Basic struct {
	client PasswordClient
	log    *slog.Logger
}

func ProviderBasic(client PasswordClient) Client {
	return &Basic{
		client: client,
		log:    log.New(ClientBasic),
	}
}

func (c *Basic) Authenticate(ctx context.Context, req *AuthRequest) (*Identity, error) {
	username, password, ok := req.HttpReq.BasicAuth()
	if !ok {
		return nil, ErrBasic
	}
	return c.client.AuthenticatePasswd(ctx, username, password)
}

func (c *Basic) Name() string {
	return ClientBasic
}

func (c *Basic) IsEnable() bool {
	return true
}
