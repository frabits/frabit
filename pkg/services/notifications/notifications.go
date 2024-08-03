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

package notifications

import (
	"context"
	"log/slog"

	"github.com/frabits/frabit/pkg/infra/log"
)

type Service struct {
	EmailQueen   chan interface{}
	WebHookQueen chan interface{}
	Logger       *slog.Logger
}

func ProviderService() *Service {
	ds := &Service{
		Logger: log.New("notifications"),
	}

	return ds
}

func (s *Service) Run(ctx context.Context) error {
	s.Logger.Info("Notification service started")
	for {
		select {
		case <-ctx.Done():
			s.Logger.Info("Notification service shutdown")
			return ctx.Err()
		case Email := <-s.EmailQueen:
			s.Logger.Info("Send email", "Email", Email)
			return nil
		case WebHook := <-s.WebHookQueen:
			s.Logger.Info("Execute webHook", "WebHook", WebHook)
			return nil
		}
	}
}
