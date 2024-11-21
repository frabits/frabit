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

package session

import (
	"context"
)

type Service interface {
	CreateSession(ctx context.Context, auth *CreateSessionCmd) (string, error)
	LookupSession(ctx context.Context, unhashedToken string) (*Session, error)
	TryRotateSession(ctx context.Context, token *Session, auth *CreateSessionCmd) (bool, *Session, error)
	RevokeSession(ctx context.Context, token *Session, soft bool) error
	RevokeAllUserSessions(ctx context.Context, login string) error
	GetUserSession(ctx context.Context, userId, userTokenId int64) (*Session, error)
	GetUserSessions(ctx context.Context, userId int64) ([]*Session, error)
	GetUserRevokedSessions(ctx context.Context, userId int64) ([]*Session, error)
}
