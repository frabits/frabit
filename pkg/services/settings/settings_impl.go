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

import (
	"context"
	"encoding/base64"
	"log/slog"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/secrets"
)

type service struct {
	secrets secrets.Service
	log     *slog.Logger
	store   Store
}

func ProviderService(secrets secrets.Service, meta *db.MetaStore) Service {
	metaStore := newStoreImpl(meta.Gdb)
	return &service{
		secrets: secrets,
		log:     log.New("settings"),
		store:   metaStore,
	}
}

func (s *service) CreateSettings(ctx context.Context, cmd *CreateSettingsCmd) error {
	encSettings, err := s.secrets.Encrypt([]byte(cmd.Settings))
	if err != nil {
		s.log.Error("create sso settings failed", "Error", err.Error())
		return err
	}
	settings := &SettingsSSO{
		Name:      cmd.Name,
		Settings:  base64.RawStdEncoding.EncodeToString(encSettings),
		CreatedAt: utils.NowDatetime(),
		UpdatedAt: utils.NowDatetime(),
	}

	return s.store.CreateSettings(ctx, settings)
}

func (s *service) GetSettings(ctx context.Context) ([]SettingsSSO, error) {
	results := make([]SettingsSSO, 0)
	settings, err := s.store.QuerySettings(ctx)
	if err != nil {
		s.log.Error("query settings failed", "Error", err.Error())
	}
	for _, setting := range settings {
		result := SettingsSSO{
			Id:        setting.Id,
			Name:      setting.Name,
			CreatedAt: setting.CreatedAt,
			UpdatedAt: setting.UpdatedAt,
		}
		toDecrypt, _ := base64.RawStdEncoding.DecodeString(setting.Settings)
		decryptSetting, err := s.secrets.Decrypt(toDecrypt)
		if err != nil {
			s.log.Error("decrypt settings failed", "Error", err.Error())
		}
		result.Settings = string(decryptSetting)

		results = append(results, result)
	}
	return results, nil
}
