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

package cleanup

import (
	"context"
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/infra/log"
)

type Service struct {
	Id     int
	Logger *slog.Logger
}

func ProviderService() *Service {
	ds := &Service{
		Logger: log.New("cleanup"),
	}

	return ds
}

func (ds *Service) Run(ctx context.Context) error {
	ds.Logger.Info("Cleanup service start")
	tick := time.NewTicker(60 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			ds.Logger.Info("shutdown service successfully")
			return ctx.Err()
		case <-tick.C:
			ds.Logger.Info("Im working now")
		}
	}
}
