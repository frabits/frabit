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

package apikey

import (
	"context"
	"errors"
)

var (
	ErrApikeyNotExists = errors.New("api key not exists")
)

type Service interface {
	AddAPIKey(ctx context.Context, cmd *CreateAPIKeyCmd) (*APIKey, error)
	GetAPIKeyByHash(ctx context.Context, hash string) (*APIKey, error)
}
