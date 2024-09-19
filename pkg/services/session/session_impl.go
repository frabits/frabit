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

package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
)

type service struct {
	cfg   *config.Config
	store store
	log   *slog.Logger
}

func ProviderService(cfg *config.Config, metaDB *db.MetaStore) Service {
	metaStore := NewStoreImpl(metaDB.Gdb)
	return &service{
		cfg:   cfg,
		store: metaStore,
		log:   log.New("auth"),
	}
}

func (s *service) CreateToken(ctx context.Context, auth *CreateUserAuth) (string, error) {
	token, hashedToken, _ := generateToken()
	userAuth := &UserAuth{
		Login:     auth.Login,
		ClientIP:  auth.ClientIP,
		UserAgent: auth.UserAgent,
		PrevToken: hashedToken,
		Token:     hashedToken,
		CreatedAt: utils.NowDatetime(),
		UpdatedAt: utils.NowDatetime(),
	}
	err := s.store.CreateToken(ctx, userAuth)
	if err != nil {
		s.log.Error("create user token failed", "Error", err.Error())
		return "", err
	}
	return token, nil
}

// LookupToken query user session token via unhashedToken
func (s *service) LookupToken(ctx context.Context, unhashedToken string) (*UserAuth, error) {
	userAuth := &UserAuth{}
	return userAuth, nil
}

func (s *service) TryRotateToken(ctx context.Context, token *UserAuth, auth *CreateUserAuth) (bool, *UserAuth, error) {
	userAuth := &UserAuth{}
	return true, userAuth, nil
}

func (s *service) RevokeToken(ctx context.Context, token *UserAuth, soft bool) error {
	return nil
}

func (s *service) RevokeAllUserTokens(ctx context.Context, userId int64) error {
	return nil
}

func (s *service) GetUserToken(ctx context.Context, userId, userTokenId int64) (*UserAuth, error) {
	userAuth := &UserAuth{}
	return userAuth, nil
}

func (s *service) GetUserTokens(ctx context.Context, userId int64) ([]*UserAuth, error) {
	userAuths := []*UserAuth{}
	return userAuths, nil
}

func (s *service) GetUserRevokedTokens(ctx context.Context, userId int64) ([]*UserAuth, error) {
	userAuths := []*UserAuth{}
	return userAuths, nil
}

func createToken() string {
	token := utils.GenRandom(32)
	return token
}

func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token + config.Security))
	return hex.EncodeToString(h.Sum(nil))
}

func generateToken() (string, string, error) {
	token := createToken()
	return token, hashToken(token), nil
}
