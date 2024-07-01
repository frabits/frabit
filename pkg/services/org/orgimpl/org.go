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

package orgimpl

import (
	"context"
	"log/slog"

	fblog "github.com/frabits/frabit/pkg/log"
	"github.com/frabits/frabit/pkg/services/org"
)

type service struct {
	store  store
	logger *slog.Logger
}

func NewService() *service {
	log := fblog.New("org")
	return &service{
		logger: log,
	}
}

func (s *service) CreateOrg(ctx context.Context, org *org.Org) (int64, error) {
	return s.store.Insert(ctx, org)
}
