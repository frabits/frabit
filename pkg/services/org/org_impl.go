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

package org

import (
	"context"
	"log/slog"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"

	fb "github.com/frabits/frabit-go-sdk/frabit"
)

type service struct {
	cfg   *config.Config
	store store
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, metaDB *db.MetaStore) Service {
	metaStore := NewStoreImpl(metaDB.Gdb)
	return &service{
		cfg:   cfg,
		store: metaStore,
		log:   log.New("org"),
	}
}

func (s *service) GetOrgs(ctx context.Context) ([]Org, error) {
	return s.store.GetOrgs(ctx)
}

func (s *service) GetOrgByName(ctx context.Context, name string) (Org, error) {
	return s.store.GetOrgByName(ctx, name)
}

func (s *service) CreateOrg(ctx context.Context, req fb.OrgCreateRequest) (int64, error) {
	s.log.Info("")
	org := &Org{
		Name:        req.Name,
		Description: req.Description,
		Country:     req.Country,
		CreatedAt:   utils.NowDatetime(),
		UpdatedAt:   utils.NowDatetime(),
	}
	gid, err := s.store.CreateOrg(ctx, org)
	if err != nil {
		s.log.Error("create org failed", "Error", err.Error())
		return 0, err
	}
	s.log.Info("create org successfully", "gid", gid)
	return gid, nil
}

func (s *service) UpdateOrg(ctx context.Context, req fb.OrgUpdateRequest) error {
	s.log.Info("")
	org := &Org{
		Name:        req.Name,
		Description: req.Description,
		Country:     req.Country,
		UpdatedAt:   utils.NowDatetime(),
	}
	err := s.store.Update(ctx, org)
	if err != nil {
		s.log.Error("create org failed", "Error", err.Error())
		return err
	}
	s.log.Info("update org successfully")
	return nil
}
