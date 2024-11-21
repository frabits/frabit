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

package authn

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	ClientAPIKey = "authn.client.apikey"
	ClientBasic  = "authn.client.basic"
	ClientForm   = "authn.client.form"

	BasicPrefix  = "Basic "
	BearerPrefix = "Bearer "
)

var (
	ErrNotFoundToken      = errors.New("not found token from current request")
	ErrUserNotExists      = errors.New("username not exists")
	ErrCredential         = errors.New("username or password not correct")
	ErrFailedLoginTooMuch = errors.New("failed login too much")
)

type Authenticator interface {
	Authenticate(ctx context.Context, req *AuthRequest) (*Identity, error)
}

type Client interface {
	Authenticator
	Name() string
	IsEnable() bool
}

type PasswordClient interface {
	AuthenticatePasswd(ctx context.Context, req *http.Request, username string, password string) (*Identity, error)
}

func ClientWithPrefix(name string) string {
	return fmt.Sprintf("authn.client.%s", name)
}

func GetTokenFromRequest(req *http.Request) (string, error) {
	header := req.Header.Get("Authorization")
	if strings.HasPrefix(header, BearerPrefix) {
		return strings.TrimPrefix(header, BearerPrefix), nil
	}
	if strings.HasPrefix(header, BasicPrefix) {
		return strings.TrimPrefix(header, BasicPrefix), nil
	}
	return "", ErrNotFoundToken
}
