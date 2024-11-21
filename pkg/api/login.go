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
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/frabits/frabit/pkg/services/authn"
	"github.com/frabits/frabit/pkg/services/login"
	"github.com/frabits/frabit/pkg/services/session"

	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) Login(c *gin.Context) {
	// 1.native
	// 2.ldap
	var authInfo authn.AuthRequest
	if err := c.ShouldBindJSON(&authInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}
	if err := hs.remoteCache.Set(hs.ctx, "native_login"+authInfo.Username, []byte(authInfo.Password), 10*time.Minute); err != nil {
		hs.Logger.Error("create remoteCache failed", "Error", err.Error())
	}
	authReq := &authn.AuthRequest{
		HttpReq:  c.Request,
		Username: authInfo.Username,
		Password: authInfo.Password,
	}

	id, err := hs.authn.Login(hs.ctx, authn.ClientForm, authReq)
	if err != nil {
		hs.Logger.Error("Login failed", "Error", err.Error())
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
	} else {
		// create a session token
		// hs.Logger.Info("get identity", "Identity", id)
		userSession := &session.CreateSessionCmd{
			Login:     id.Login,
			ClientIP:  c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
		}
		sess, err := hs.session.CreateSession(hs.ctx, userSession)
		if err != nil {
			hs.Logger.Error("create session token failed", "Error", err.Error())
		}
		encSession, err := hs.secrets.Encrypt([]byte(sess))
		c.Request.Header.Set("frabit", sess)
		result := login.LoginDTO{
			Msg:   "Login successfully",
			Token: hex.EncodeToString(encSession),
		}
		c.IndentedJSON(http.StatusOK, result)
	}
}

func (hs *HttpServer) LoginOauth(c *gin.Context) {
	// 1.SSO: github、google、OIDC
	client := c.Param("name")
	authReq := &authn.AuthRequest{
		HttpReq: c.Request,
	}
	id, err := hs.authn.Login(hs.ctx, authn.ClientWithPrefix(client), authReq)
	if err != nil {
		hs.Logger.Error("Login failed", "Error", err.Error())
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
	} else {
		// create a session token
		userSession := &session.CreateSessionCmd{
			Login:     id.Login,
			ClientIP:  c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
		}
		sess, err := hs.session.CreateSession(hs.ctx, userSession)
		if err != nil {
			hs.Logger.Error("create session token failed", "Error", err.Error())
		}
		encSession, err := hs.secrets.Encrypt([]byte(sess))
		c.Request.Header.Set("frabit", sess)
		result := login.LoginDTO{
			Msg:   "Login successfully",
			Token: base64.StdEncoding.EncodeToString(encSession),
		}
		c.IndentedJSON(http.StatusOK, result)
	}
}

func (hs *HttpServer) Logout(c *gin.Context) {
	// remove session
	// redirect to login page
	token, err := authn.GetTokenFromRequest(c.Request)
	if err != nil {
		hs.Logger.Error("not found session from header", "Error", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
	}
	redirect, err := hs.authn.Logout(hs.ctx, token)
	if err != nil {
		hs.Logger.Error("Login failed", "Error", err.Error())
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
	}
	c.Redirect(http.StatusPermanentRedirect, redirect.URL)
}
