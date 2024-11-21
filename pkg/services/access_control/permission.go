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

package access_control

import (
	"context"
	"log/slog"

	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
)

type permissionImpl struct {
	log   *slog.Logger
	store Store
}

func ProviderService(db *db.MetaStore) Service {
	store := providerStore(db.Gdb)
	pi := &permissionImpl{
		log:   log.New("permission"),
		store: store,
	}
	return pi
}

func (s *permissionImpl) CreatePermission(ctx context.Context) error {
	return nil
}
