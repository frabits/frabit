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

import (
	"context"
	"gorm.io/gorm"
)

type Store interface {
	CreateTask(ctx context.Context, task *Task) error
	GetTasks(ctx context.Context) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) ([]*Task, error)
}

type storeImpl struct {
	db *gorm.DB
}

func providerStore(db *gorm.DB) Store {
	return &storeImpl{db: db}
}

func (s *storeImpl) CreateTask(ctx context.Context, task *Task) error {
	return nil
}

func (s *storeImpl) GetTasks(ctx context.Context) ([]*Task, error) {
	return nil, nil
}

func (s *storeImpl) UpdateTask(ctx context.Context, task *Task) ([]*Task, error) {
	return nil, nil
}