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

package registry

import "context"

// BackgroundServiceRegistry provides background services
type BackgroundServiceRegistry interface {
	GetServices() []BackgroundService
}

type BackgroundService interface {
	// Run start the background process of the service after `Init` have been revoked on
	// all service. The context.Context passed into function should be used to subscribe to ctx.Done
	// so that the service can be notified when Frabit shuts down.
	Run(ctx context.Context) error
}
