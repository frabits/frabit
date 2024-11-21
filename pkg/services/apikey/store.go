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
	"github.com/frabits/frabit/pkg/common/utils"

	"gorm.io/gorm"
)

type Store interface {
	AddAPIKey(ctx context.Context, key *APIKey) error
	GetAPIKeyByHash(ctx context.Context, hash string) (*APIKey, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) Store {
	return &storeImpl{DB: db}
}

func (s *storeImpl) AddAPIKey(ctx context.Context, key *APIKey) error {
	s.DB.Create(key)
	return nil
}

func (s *storeImpl) GetAPIKeyByHash(ctx context.Context, hash string) (*APIKey, error) {
	var apikey APIKey
	s.DB.Where("hash_key=?", hash).First(&apikey)
	s.updateLastSeen(ctx, apikey.Id)
	return &apikey, nil
}

func (s *storeImpl) updateLastSeen(ctx context.Context, keyId uint32) error {
	s.DB.Model(&APIKey{}).Where("id=?", keyId).Update("last_used_at", utils.NowDatetime())
	return nil
}
