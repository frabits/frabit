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

package serviceaccount

import (
	"context"
	"fmt"
	fb "github.com/frabits/frabit-go-sdk/frabit"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"

	"github.com/frabits/frabit/pkg/services/apikey"
	"github.com/frabits/frabit/pkg/services/org"
	"github.com/frabits/frabit/pkg/services/user"
)

type serviceAccountImpl struct {
	apiKeySrc apikey.Service
	userSrv   user.Service
	orgSrv    org.Service
	log       *slog.Logger
}

func ProviderService(apikey apikey.Service, user user.Service, org org.Service) Service {
	sa := &serviceAccountImpl{
		apiKeySrc: apikey,
		userSrv:   user,
		orgSrv:    org,
		log:       log.New("service.account"),
	}

	return sa
}

// CreateServiceAccount create a service account
func (s *serviceAccountImpl) CreateServiceAccount(ctx context.Context, cmd *CreateServiceAccountCmd) error {
	user := fb.CreateUserRequest{
		Login:            cmd.Name,
		Name:             fmt.Sprintf("sa_%s", utils.CreateUUIDWithDelimiter("")),
		IsServiceAccount: 1,
	}
	uid, err := s.userSrv.CreateUser(ctx, user)
	if err != nil {
		s.log.Error("create service account failed", "Error", err.Error())
		return err
	}
	s.log.Info("create service account successfully", "ServiceAccountId", uid)
	return nil
}

// DeleteServiceAccount delete specific service account and bind apikey
func (s *serviceAccountImpl) DeleteServiceAccount(ctx context.Context, name string) error {
	return nil
}

func (s *serviceAccountImpl) UpdateServiceAccount(ctx context.Context, cmd *UpdateServiceAccountCmd) error {
	return nil
}

func (s *serviceAccountImpl) DisableServiceAccount(ctx context.Context, name string) error {
	return nil
}

func (s *serviceAccountImpl) GetServiceAccount(ctx context.Context) ([]*ServiceAccountDTO, error) {
	var saResult []*ServiceAccountDTO
	saAccounts, err := s.userSrv.GetServiceAccount(ctx)
	if err != nil {
		s.log.Error("query service account failed", "Error", err.Error())
	}
	for _, account := range saAccounts {
		sa := &ServiceAccountDTO{
			OrgId:     account.OrgId,
			Name:      account.Login,
			Role:      "",
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}
		saResult = append(saResult, sa)
	}
	return saResult, nil
}

func (s *serviceAccountImpl) GetServiceAccountByName(ctx context.Context, name string) ([]*ServiceAccountDTO, error) {
	var saResult []*ServiceAccountDTO
	saAccount, err := s.userSrv.GetUserByLogin(ctx, name)
	if err != nil {
		s.log.Error("query service account failed", "Error", err.Error())
	}
	if saAccount.IsServiceAccount == 1 {
		sa := &ServiceAccountDTO{
			OrgId:     saAccount.OrgId,
			Name:      saAccount.Login,
			Role:      "",
			CreatedAt: saAccount.CreatedAt,
			UpdatedAt: saAccount.UpdatedAt,
		}
		saResult = append(saResult, sa)
	}
	return saResult, nil
}
