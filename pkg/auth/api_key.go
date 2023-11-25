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

package auth

import (
	"fmt"
	"github.com/frabits/frabit/pkg/common/utils"
	"time"
)

const apiKeyLength int = 32

type APIKey struct {
	Prefix      string
	PublicAuth  string
	PrivateAuth string
	VerifyAuth  string
	CreateAt    time.Time
	LastSeen    time.Time
}

func NewAPIKey() *APIKey {
	pubAuth := utils.NewToken(apiKeyLength).Hash
	priAuth := utils.NewToken(apiKeyLength).Hash
	return &APIKey{
		Prefix:      "frabit_tkn",
		PublicAuth:  pubAuth,
		PrivateAuth: priAuth,
		VerifyAuth:  utils.GenHash(pubAuth, priAuth),
		CreateAt:    time.Now(),
		LastSeen:    time.Now(),
	}
}

func (key *APIKey) IsValid(pubAuth, priAuth string) bool {
	if ok := key.VerifyAuth == utils.GenHash(pubAuth, priAuth); ok {
		return true
	}
	return false
}

func (key *APIKey) PublicKey() string {
	return fmt.Sprintf("%s_%s", key.Prefix, key.PublicAuth)
}
