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

import (
	"context"
	"github.com/pkg/errors"

	fb "github.com/frabits/frabit-go-sdk/frabit"
)

type Service interface {
	CreateUser(ctx context.Context, createReq fb.CreateUserRequest) (uint32, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id uint32) (UserProfileDTO, error)
	GetServiceAccount(ctx context.Context) ([]UserProfileDTO, error)
	DeleteUser(ctx context.Context, login string) error
	GenerateCode(ctx context.Context, login string) (string, error)
	ValidateCode(ctx context.Context, code string) ([]string, error)
	UpdateUserLastSeen(ctx context.Context, login string) error
	VerifyEmail(ctx context.Context, login string) error
}

const (
	minPasswordLength = 8
)

var (
	ErrUserNotExists                   = errors.New("user not exists")
	ErrVerifyCodeInvalid               = errors.New("verify code invalid")
	ErrVerifyCodeExpired               = errors.New("verify code expired")
	ErrPasswordTooShort                = errors.New("password too short")
	ErrPasswordTooShortForStrongPolicy = errors.New("password too short for strong policy")
	ErrPasswordNotMatchStrongPolicy    = errors.New("password not match strong policy")
)
