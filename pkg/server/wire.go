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

//go:build wireinject
// +build wireinject

package server

import (
	"github.com/frabits/frabit/pkg/api"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	bgSrv "github.com/frabits/frabit/pkg/server/bg_services"
	"github.com/frabits/frabit/pkg/services/agent"
	"github.com/frabits/frabit/pkg/services/audit"
	"github.com/frabits/frabit/pkg/services/auth"
	"github.com/frabits/frabit/pkg/services/backup"
	"github.com/frabits/frabit/pkg/services/cleanup"
	"github.com/frabits/frabit/pkg/services/deploy"
	"github.com/frabits/frabit/pkg/services/license"
	"github.com/frabits/frabit/pkg/services/login"
	ns "github.com/frabits/frabit/pkg/services/notifications"
	"github.com/frabits/frabit/pkg/services/org"
	"github.com/frabits/frabit/pkg/services/secrets"
	"github.com/frabits/frabit/pkg/services/team"
	uc "github.com/frabits/frabit/pkg/services/updatechecker"
	"github.com/frabits/frabit/pkg/services/user"

	"github.com/google/wire"
)

var wireSet = wire.NewSet(
	db.New,
	config.ProviderConfig,
	agent.ProviderAgentService,
	auth.ProviderService,
	audit.ProviderService,
	backup.ProviderMySQLBackup,
	cleanup.ProviderService,
	deploy.ProviderService,
	login.ProviderLoginNative,
	license.ProviderService,
	secrets.ProviderSecrets,
	ns.ProviderService,
	org.ProviderService,
	team.ProviderService,
	user.ProviderService,
	uc.ProviderFrabitService,
	bgSrv.ProviderBackgroundServiceRegistry,
	api.ProviderHTTPServer,
	NewServer,
)

func Initialize() (*Server, error) {
	wire.Build(wireSet)
	return &Server{}, nil
}
