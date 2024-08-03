// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
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

package bg_services

import (
	"github.com/frabits/frabit/pkg/registry"
	"github.com/frabits/frabit/pkg/services/cleanup"
	"github.com/frabits/frabit/pkg/services/deploy"
	"github.com/frabits/frabit/pkg/services/license"
	ns "github.com/frabits/frabit/pkg/services/notifications"
	uc "github.com/frabits/frabit/pkg/services/updatechecker"
)

type BackgroundServiceRegistry struct {
	services []registry.BackgroundService
}

func ProviderBackgroundServiceRegistry(cleanup *cleanup.Service, notifications *ns.Service, deploy deploy.Service,
	updateChecker *uc.FrabitService, license license.Service) registry.BackgroundServiceRegistry {
	return NewBackgroundServiceRegistry(cleanup, notifications, deploy, updateChecker, license)
}

func NewBackgroundServiceRegistry(services ...registry.BackgroundService) registry.BackgroundServiceRegistry {
	return &BackgroundServiceRegistry{services}
}

func (bsr *BackgroundServiceRegistry) GetServices() []registry.BackgroundService {
	return bsr.services
}
