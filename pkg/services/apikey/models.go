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

import "errors"

var (
	ErrNotFound          = errors.New("API key not found")
	ErrGenerateFailed    = errors.New("API key not generate")
	ErrInvalid           = errors.New("invalid API key")
	ErrInvalidExpiration = errors.New("negative value for SecondsToLive")
	ErrDuplicate         = errors.New("API key, organization ID and name must be unique")
)

type APIKey struct {
	Id               uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name             string `gorm:"type:varchar(100);not null;unique:uniq_name" json:"name"`
	OrgId            uint32 `gorm:"type:varchar(200);not null" json:"org_id"`
	Role             string `gorm:"type:varchar(200);not null" json:"role"`
	HashKey          string `gorm:"type:varchar(20);not null" json:"hash_key"`
	CreatedAt        string `gorm:"type:varchar(50);not null" json:"created_at"`
	UpdatedAt        string `gorm:"type:varchar(50);not null" json:"updated_at"`
	LastUsedAt       string `gorm:"type:varchar(50);not null" json:"last_used_at"`
	Expires          string `gorm:"type:varchar(50);not null" json:"expires"`
	ServiceAccountId int32  `gorm:"type:bigint(20);not null" json:"service_account_id"`
	IsRevoked        bool   `gorm:"type:tinyint(1);not null" json:"is_revoked"`
}

func (a APIKey) TableName() string {
	return "apikey"
}

type CreateAPIKeyCmd struct {
	Name             string `json:"name"`
	OrgId            uint32 `json:"org_id"`
	Role             string `json:"role"`
	HashKey          string `json:"hash_key"`
	SecondsToLive    uint64 `json:"seconds_to_live"`
	ServiceAccountId int32  `json:"service_account_id"`
}

type APIKeyDTO struct {
	Name         string `json:"name"`
	ClientSecret string `json:"client_secret"`
}
