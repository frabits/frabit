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

package session

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
		log:   log.New("session"),
	}
}

func (s *service) CreateSession(ctx context.Context, auth *CreateSessionCmd) (string, error) {
	token, hashedToken, _ := generateToken()
	userAuth := &Session{
		Login:     auth.Login,
		ClientIP:  auth.ClientIP,
		UserAgent: auth.UserAgent,
		PrevToken: hashedToken,
		Token:     hashedToken,
		CreatedAt: utils.NowDatetime(),
		UpdatedAt: utils.NowDatetime(),
	}
	err := s.store.CreateSession(ctx, userAuth)
	if err != nil {
		s.log.Error("create user token failed", "Error", err.Error())
		return "", err
	}
	return token, nil
}

// LookupSession LookupToken query user session token via unhashedToken
func (s *service) LookupSession(ctx context.Context, unhashedToken string) (*Session, error) {
	hashedToken := hashToken(unhashedToken)
	s.log.Info("generate hashed token", "Token", hashedToken)
	session, err := s.store.GetSessionByToken(ctx, hashedToken)
	if err != nil {
		s.log.Error("not find specific session", "Error", err.Error())
		return nil, err
	}
	return session, nil
}

func (s *service) TryRotateSession(ctx context.Context, token *Session, auth *CreateSessionCmd) (bool, *Session, error) {
	userAuth := &Session{}
	return true, userAuth, nil
}

func (s *service) RevokeSession(ctx context.Context, token *Session, soft bool) error {
	if !soft {
		return s.store.DeleteSessionById(ctx, token.Id)
	}
	return nil
}

func (s *service) RevokeAllUserSessions(ctx context.Context, login string) error {
	s.log.Info("revoke all session", "Identity", login)
	return s.store.RevokeUserAllSessions(ctx, login)
}

func (s *service) GetUserSession(ctx context.Context, userId, userTokenId int64) (*Session, error) {
	userAuth := &Session{}
	return userAuth, nil
}

func (s *service) GetUserSessions(ctx context.Context, userId int64) ([]*Session, error) {
	userAuths := []*Session{}
	return userAuths, nil
}

func (s *service) GetUserRevokedSessions(ctx context.Context, userId int64) ([]*Session, error) {
	userAuths := []*Session{}
	return userAuths, nil
}

func createToken() string {
	token := utils.GenRandom(32)
	return token
}

func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

func generateToken() (string, string, error) {
	token := createToken()
	return token, hashToken(token), nil
}
