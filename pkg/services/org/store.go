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

package org

import (
	"context"
	"gorm.io/gorm"
)

type store interface {
	GetOrgs(context.Context) ([]Org, error)
	GetOrgByName(context.Context, string) (Org, error)
	Update(context.Context, *Org) error
	CreateOrg(context.Context, *Org) (int64, error)
}

type storeImpl struct {
	DB *gorm.DB
}

func NewStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{db}
}

func (s *storeImpl) GetOrgs(ctx context.Context) ([]Org, error) {
	var orgs []Org
	s.DB.Model(Org{}).Find(&orgs)
	return orgs, nil
}

func (s *storeImpl) GetOrgByName(ctx context.Context, name string) (Org, error) {
	var org Org
	s.DB.Model(Org{}).Where("name = ?", name).First(&org)
	return org, nil
}

func (s *storeImpl) Update(ctx context.Context, org *Org) error {
	s.DB.Create(org)
	return nil
}

func (s *storeImpl) CreateOrg(ctx context.Context, org *Org) (int64, error) {
	s.DB.Create(org)
	return 0, nil
}
