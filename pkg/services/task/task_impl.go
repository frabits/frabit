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

package task

import (
	"context"
	"log/slog"

	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/secrets"
)

type taskImpl struct {
	log    *slog.Logger
	secret secrets.Service
	store  Store
}

func ProviderService(secret secrets.Service, metaDB *db.MetaStore) Service {
	store := providerStore(metaDB.Gdb)
	ti := &taskImpl{
		log:    log.New("task"),
		secret: secret,
		store:  store,
	}

	return ti
}

func (s *taskImpl) AddTask(ctx context.Context, cmd *CreateTaskCmd) error {
	return nil
}

func (s *taskImpl) UpdateTask(ctx context.Context, cmd *UpdateTaskCmd) error {
	return nil
}

func (s *taskImpl) ListTasks(ctx context.Context) ([]*Task, error) {
	return nil, nil
}
