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

package task

type Task struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	AgentId   uint32 `gorm:"type:varchar(50);not null;unique:uniq_name" json:"agent_id"`
	Name      string `gorm:"type:varchar(50);not null;unique:uniq_name" json:"name"`
	TaskType  string `gorm:"type:varchar(200);not null" json:"description"`
	State     string `gorm:"type:varchar(30);not null" json:"state"`
	Creator   string `gorm:"type:varchar(30);not null" json:"creator"`
	CreatedAt int64  `gorm:"type:bigint(22);not null" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint(22);not null" json:"updated_at"`
}

func (a Task) TableName() string {
	return "task"
}

type CreateTaskCmd struct {
	Name     string `json:"name"`
	TaskType string `json:"description"`
	Owner    string `json:"owner"`
}

type UpdateTaskCmd struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
