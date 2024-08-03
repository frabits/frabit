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

package auth

import (
	"context"
	"time"
)

func (s *service) Run(ctx context.Context) error {
	s.log.Info("Notification service started")
	ticker := time.NewTicker(time.Hour)
	maxInactiveLifetime := s.cfg.Server.LoginMaxInactiveLifetime
	maxLifetime := s.cfg.Server.LoginMaxLifetime
	for {
		select {
		case <-ticker.C:
			s.log.Info("Execute webHook", "WebHook")
			if err := s.deleteExpiredTokens(ctx, maxInactiveLifetime, maxLifetime); err != nil {
				s.log.Error("An error occurred while deleting expired tokens", "err", err)
			}
		case <-ctx.Done():
			s.log.Info("Auth cleanup service shutdown")
			return ctx.Err()
		}
	}
}

func (s *service) deleteExpiredTokens(ctx context.Context, maxInactiveLifetime, maxLifetime time.Duration) error {
	createdBefore := time.Now().Add(-maxLifetime)
	rotatedBefore := time.Now().Add(-maxInactiveLifetime)

	s.log.Debug("starting cleanup of expired auth tokens", "createdBefore", createdBefore, "rotatedBefore", rotatedBefore)
	return s.store.DeleteTokenByTime(ctx, createdBefore, rotatedBefore)
}
