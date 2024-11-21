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

package apikey

import (
	"context"
	"github.com/frabits/frabit/pkg/common/constant"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/satoken"
	"github.com/frabits/frabit/pkg/services/user"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
)

type serviceImpl struct {
	store Store
	user  user.Service
	cfg   *config.Config
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, meta *db.MetaStore, user user.Service) Service {
	metaStore := NewStoreImpl(meta.Gdb)
	si := &serviceImpl{
		store: metaStore,
		user:  user,
		cfg:   cfg,
		log:   log.New("apikey"),
	}
	return si
}

func (s *serviceImpl) AddAPIKey(ctx context.Context, cmd *CreateAPIKeyCmd) (*APIKey, error) {
	Expires := ""
	if cmd.SecondsToLive < 0 {
		s.log.Error(ErrInvalidExpiration.Error())
		return nil, ErrInvalidExpiration
	} else if cmd.SecondsToLive > 0 {
		Expires = time.Now().Add(time.Second * time.Duration(cmd.SecondsToLive)).Format(time.DateTime)
	}

	apiKey := &APIKey{
		Name:             cmd.Name,
		OrgId:            constant.GlobalOrgId,
		Role:             cmd.Role,
		HashKey:          cmd.HashKey,
		CreatedAt:        utils.NowDatetime(),
		UpdatedAt:        utils.NowDatetime(),
		LastUsedAt:       utils.NowDatetime(),
		Expires:          Expires,
		ServiceAccountId: cmd.ServiceAccountId,
	}
	if err := s.store.AddAPIKey(ctx, apiKey); err != nil {
		s.log.Error("create apikey failed", "Error", err.Error())
		return nil, ErrGenerateFailed
	}
	return apiKey, nil
}

func (s *serviceImpl) GetAPIKeyByHash(ctx context.Context, hash string) (*APIKey, error) {
	prefixToken, err := satoken.Decode(hash)
	if err != nil {
		s.log.Error("cannot decode service account token", "Error", err.Error())
		return nil, err
	}
	s.log.Info("decode token from header", "token", prefixToken)

	return s.store.GetAPIKeyByHash(ctx, prefixToken.Hash())
}
