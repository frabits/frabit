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

package settings

import (
	"context"
	"gorm.io/gorm"
)

type Store interface {
	CreateSettings(ctx context.Context, settings *SettingsSSO) error
	QuerySettings(ctx context.Context) ([]SettingsSSO, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func newStoreImpl(db *gorm.DB) Store {
	return &storeImpl{DB: db}
}

func (s *storeImpl) CreateSettings(ctx context.Context, settings *SettingsSSO) error {
	s.DB.Create(settings)
	return nil
}

func (s *storeImpl) QuerySettings(ctx context.Context) ([]SettingsSSO, error) {
	var settings []SettingsSSO
	s.DB.Model(SettingsSSO{}).Find(&settings)
	return settings, nil
}
