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

	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/role"
	"github.com/frabits/frabit/pkg/services/user"
)

type AccessControl interface {
	Evaluate(ctx context.Context, user string, eval Evaluator) (bool, error)
}

type Service interface {
	CreatePermission(ctx context.Context) error
}

type accessControlImpl struct {
	log        *slog.Logger
	userSrv    user.Service
	roleSrv    role.Service
	permission Service
}

func ProviderAccessControl(userSrv user.Service, roleSrv role.Service, permission Service) AccessControl {
	aci := &accessControlImpl{
		log: log.New("access.control"),
	}
	aci.userSrv = userSrv
	aci.roleSrv = roleSrv
	aci.permission = permission
	return aci
}

func (s *accessControlImpl) Evaluate(ctx context.Context, user string, eval Evaluator) (bool, error) {
	// ger user permission

	return true, nil
}
