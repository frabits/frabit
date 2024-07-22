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

package agent

type Agent struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Hostname  string `gorm:"type:varchar(100);not null;unique:uniq_hostname" json:"hostname"`
	AgentId   string `gorm:"type:varchar(50);not null;unique:uniq_agent_id" json:"agent_id"`
	Label     string `gorm:"type:varchar(30);not null" json:"label"`
	ClientIp  string `gorm:"type:varchar(50);not null" json:"client_ip"`
	CreatedAt string `gorm:"type:varchar(50);not null" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(50);not null" json:"updated_at"`
}

func (a Agent) TableName() string {
	return "agent"
}

type Heartbeat struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	AgentId   string `gorm:"type:varchar(50);not null;unique:uniq_agent_id" json:"agent_id"`
	Status    string `gorm:"type:varchar(30);not null" json:"status"`
	CreatedAt string `gorm:"type:varchar(50);not null" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(50);not null" json:"updated_at"`
}

func (a Heartbeat) TableName() string {
	return "agent_heartbeat"
}
