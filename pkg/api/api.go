// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
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
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/frabits/frabit/pkg/satoken"
	ac "github.com/frabits/frabit/pkg/services/access_control"
	"github.com/frabits/frabit/pkg/services/authn"
	"github.com/gin-gonic/gin"
)

// setup register all router
func (hs *HttpServer) setup(engine *gin.Engine) {
	engine.GET("/health", hs.health)
	engine.GET("/info", hs.info)
	engine.POST("/user/verify-email", hs.VerifyEmailComplete)
	engine.POST("/login", hs.Login)
	engine.POST("/login/:name", hs.LoginOauth)
	apiV2 := engine.Group("/api/v2")
	apiV2.Use(hs.auth())
	hs.applyAdmin(apiV2)
	hs.applyAudit(apiV2)
	hs.applyApiKey(apiV2)
	hs.applyOrg(apiV2)
	hs.applyTeam(apiV2)
	hs.applyUser(apiV2)
	hs.applyAgent(apiV2)
	hs.applyBackup(apiV2)
	hs.applyDeploy(apiV2)
	hs.applyServiceAccount(apiV2)
}

func (hs *HttpServer) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		var identity *authn.Identity
		token, err := authn.GetTokenFromRequest(req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}
		if strings.HasPrefix(token, satoken.FrabitPrefix+"_") {
			hs.Logger.Info("auth via service_account")
			authReq := &authn.AuthRequest{
				HttpReq: req,
			}
			identity, err = hs.authn.Login(hs.ctx, authn.ClientAPIKey, authReq)
		} else {
			hs.Logger.Info("found token", "Token", token)
			deocdeToken, _ := hex.DecodeString(token)
			decryptToken, _ := hs.secrets.Decrypt(deocdeToken)
			session, err := hs.session.LookupSession(hs.ctx, string(decryptToken))
			if err != nil {
				hs.Logger.Error("cannot get session via current token", "Error", err.Error())
			} else {
				hs.Logger.Info("generate identity from session", "Identity", session)
				identity = identity.IdentityFromSession(session)
			}
		}
		// from now on ,we should find a real identity
		if identity == nil {
			hs.Logger.Warn("not found identity")
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}
		hs.Logger.Info("Set identity", "identity", identity)
		c.Set(authn.AuthIdentity, identity.Login)
		c.Next()
	}
}

// determines the verified user's access to resources and operations
func (hs *HttpServer) authZ(evl ac.Evaluator) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.GetString(authn.AuthIdentity)
		hs.Logger.Info("Start evaluate permission", "Identity", login)
		if login == "" {
			hs.Logger.Warn("not found identity")
			c.JSON(http.StatusForbidden, "Not found identity")
			c.Abort()
			return
		}
		hasAccess, err := hs.accessControl.Evaluate(hs.ctx, login, evl)
		if err != nil {
			c.JSON(http.StatusForbidden, map[string]string{
				"Title": "Forbidden",
				"Error": err.Error(),
			})
			c.Abort()
			return
		}
		if !hasAccess {
			hs.Logger.Info("Not allowed access", "Identity", login)
			c.JSON(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}
