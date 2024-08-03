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

package user

type User struct {
	Id              uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Login           string `gorm:"type:varchar(30);not null;unique:uniq_login" json:"login"`
	Username        string `gorm:"type:varchar(200);not null;index:idx_username" json:"username"`
	Email           string `gorm:"type:varchar(100);not null;index:idx_email" json:"email"`
	Password        string `gorm:"type:varchar(200);not null" json:"password"`
	Rands           string `gorm:"type:varchar(100);not null" json:"rands"`
	IsAdmin         int    `gorm:"type:tinyint(1);not null" json:"is_admin"`
	IsDisabled      int    `gorm:"type:tinyint(1);not null" json:"is_disabled"`
	IsExternal      int    `gorm:"type:tinyint(1);not null" json:"is_external"`
	IsEmailVerified int    `gorm:"type:tinyint(1);not null" json:"is_email_verified"`
	Theme           string `gorm:"type:varchar(10);not null" json:"theme"`
	OrgId           uint32 `gorm:"type:bigint(10);not null" json:"org_id"`
	CreatedAt       string `gorm:"type:varchar(50);not null" json:"Created_at"`
	UpdatedAt       string `gorm:"type:varchar(50);not null" json:"updated_at"`
	LastSeenAt      string `gorm:"type:varchar(50);not null" json:"last_seen_at"`
}

func (u User) TableName() string {
	return "user"
}

type UserProfileDTO struct {
	Login           string `json:"login"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	IsAdmin         int    `json:"is_admin"`
	IsDisabled      int    `json:"is_disabled"`
	IsExternal      int    `json:"is_external"`
	IsEmailVerified int    `json:"is_email_verified"`
	Theme           string `json:"theme"`
	OrgId           uint32 `json:"org_id"`
	CreatedAt       string `json:"Created_at"`
	UpdatedAt       string `json:"updated_at"`
}
