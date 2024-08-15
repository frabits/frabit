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

package audit

import (
	"context"
	"github.com/frabits/frabit/pkg/bus"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"

	"github.com/frabits/frabit/pkg/config"
)

type service struct {
	cfg   *config.Config
	bus   bus.Bus
	store Store
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, meta *db.MetaStore, bus bus.Bus) Service {
	metaStore := NewStoreImpl(meta.Gdb)
	return &service{
		cfg:   cfg,
		bus:   bus,
		store: metaStore,
		log:   log.New("audit"),
	}
}

func (s *service) AddAuditEvent(ctx context.Context, cmd *CreateAuditCmd) error {
	s.log.Info("add event", "EventName", cmd.EventName)
	auditEvent := &AuditLog{
		Username:  cmd.Username,
		EventName: cmd.EventName,
		ClientIp:  cmd.ClientIp,
		CreatedAt: utils.NowDatetime(),
	}

	return s.store.Create(ctx, auditEvent)
}

func (s *service) GetAuditEvent(ctx context.Context) ([]AuditLog, error) {
	return s.store.GetEvent(ctx)
}
