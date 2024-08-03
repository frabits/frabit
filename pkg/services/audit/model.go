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

package audit

import "github.com/frabits/frabit/pkg/event"

type AuditLog struct {
	Id        uint32          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username  string          `gorm:"type:varchar(100);not null;unique:uniq_hostname" json:"username"`
	EventName event.EventType `gorm:"type:varchar(100);not null;unique:uniq_hostname" json:"event_name"`
	ClientIp  string          `gorm:"type:varchar(50);not null" json:"client_ip"`
	CreatedAt string          `gorm:"type:varchar(50);not null" json:"created_at"`
}

func (a AuditLog) TableName() string {
	return "audit_log"
}

type CreateAuditCmd struct {
	Username  string          `json:"username"`
	EventName event.EventType `json:"event_name"`
	ClientIp  string          `json:"client_ip"`
}
