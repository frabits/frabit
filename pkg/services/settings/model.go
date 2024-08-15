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

type SettingsSSO struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"type:varchar(30);not null;unique:uniq_login" json:"name"`
	Settings  string `gorm:"type:text;not null" json:"settings"`
	CreatedAt string `gorm:"type:varchar(50);not null" json:"Created_at"`
	UpdatedAt string `gorm:"type:varchar(50);not null" json:"updated_at"`
}

func (u SettingsSSO) TableName() string {
	return "settings_sso"
}

type CreateSettingsCmd struct {
	Name     string `json:"name"`
	Settings string `json:"settings"`
}
