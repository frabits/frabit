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

package team

import (
	"context"

	"gorm.io/gorm"
)

type store interface {
	Create(context.Context, *Team) (uint32, error)
	GetAll(context.Context) ([]Team, error)
	GetTeamByName(context.Context, string) (Team, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func newStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{db}
}

func (s *storeImpl) Create(ctx context.Context, team *Team) (uint32, error) {
	s.DB.Create(team)
	return 0, nil
}

func (s *storeImpl) GetAll(ctx context.Context) ([]Team, error) {
	var teams []Team
	s.DB.Model(Team{}).Find(&teams)
	return teams, nil
}

func (s *storeImpl) GetTeamByName(ctx context.Context, name string) (Team, error) {
	var team Team
	s.DB.Model(Team{}).Where("name=?", name).First(&team)
	return team, nil
}
