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

package team

import (
	"context"
	"log/slog"

	fb "github.com/frabits/frabit-go-sdk/frabit"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
)

type service struct {
	cfg   *config.Config
	store store
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, metaDB *db.MetaStore) Service {
	store := newStoreImpl(metaDB.Gdb)
	return &service{
		cfg:   cfg,
		store: store,
		log:   log.New("team"),
	}
}

func (s *service) Create(ctx context.Context, teamReq fb.CreateTeamRequest) error {
	team := &Team{
		Name:        teamReq.Name,
		Description: teamReq.Description,
		Owner:       teamReq.Owner,
		CreatedAt:   utils.NowDatetime(),
		UpdatedAt:   utils.NowDatetime(),
	}
	tid, err := s.store.Create(ctx, team)
	if err != nil {
		return err
	}
	s.log.Info("create team successfully", "team_id", tid)
	return nil
}

func (s *service) GetAll(ctx context.Context) ([]Team, error) {
	return s.store.GetAll(ctx)
}

func (s *service) GetTeamByName(ctx context.Context, name string) (Team, error) {
	return s.store.GetTeamByName(ctx, name)
}
