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

import "time"

type Login struct {
	Id            uint32    `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Rands         string    `json:"rands"`
	Email         string    `json:"email"`
	IsAdmin       bool      `json:"is_admin"`
	Disabled      bool      `json:"disable"`
	EmailVerified bool      `json:"email_verified"`
	Theme         string    `json:"theme"`
	OrgId         uint32    `json:"org_id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
	LastSeenAt    time.Time `json:"last_seen_at"`
}
