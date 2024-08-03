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

package api

import (
	"github.com/frabits/frabit/pkg/services/auth"
	"net/http"

	"github.com/frabits/frabit/pkg/services/login"

	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) Login(c *gin.Context) {
	// 1.SSO: github、google、OIDC
	// 2.native
	// 3.ldap
	var authInfo login.AuthPasswd
	if err := c.ShouldBindJSON(&authInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}

	err := hs.login.Authenticator(hs.ctx, authInfo)
	if err != nil {
		hs.Logger.Error("Login failed", "Error", err.Error())
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
	} else {
		// create a session token
		userSession := &auth.CreateUserAuth{
			Login:     authInfo.Login,
			ClientIP:  c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
		}
		session, err := hs.authUser.CreateToken(hs.ctx, userSession)
		if err != nil {
			hs.Logger.Error("create session token failed", "Error", err.Error())
		}
		c.Request.Header.Set("frabit", session)
		result := login.LoginDTO{
			Msg:   "Login successfully",
			Token: session,
		}
		c.IndentedJSON(http.StatusOK, result)
	}
}
