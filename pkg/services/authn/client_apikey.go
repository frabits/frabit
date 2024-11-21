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
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/apikey"
	"github.com/frabits/frabit/pkg/services/user"
	"log/slog"
)

type ApiKey struct {
	apikey apikey.Service
	user   user.Service
	log    *slog.Logger
}

func ProviderApiKey(apikey apikey.Service, user user.Service) Client {
	return &ApiKey{
		apikey: apikey,
		user:   user,
		log:    log.New(ClientAPIKey),
	}
}

func (c *ApiKey) Authenticate(ctx context.Context, req *AuthRequest) (*Identity, error) {
	key, err := GetTokenFromRequest(req.HttpReq)
	if err != nil {
		return nil, err
	}
	c.log.Debug("get apikey detail from header", "apikey", key)
	ret, err := c.apikey.GetAPIKeyByHash(ctx, key)
	if err != nil {
		c.log.Error("cannot get apikey", "Error")
		return nil, apikey.ErrApikeyNotExists
	}
	c.log.Debug("get apikey detail", "apikey", ret)
	// generate Identity based apikey
	u, err := c.user.GetUserById(ctx, uint32(ret.ServiceAccountId))
	if err != nil {
		return nil, user.ErrUserNotExists
	}
	id := &Identity{
		OrgId:   1,
		Name:    u.Username,
		Login:   u.Login,
		Email:   u.Email,
		IsAdmin: utils.Number2Bool(u.IsAdmin),
	}
	return id, nil
}

func (c *ApiKey) Name() string {
	return ClientAPIKey
}

func (c *ApiKey) IsEnable() bool {
	return true
}
