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
	"net/http"

	"github.com/frabits/frabit/pkg/event"
	ac "github.com/frabits/frabit/pkg/services/access_control"
	"github.com/frabits/frabit/pkg/services/audit"
	"github.com/frabits/frabit/pkg/services/authn"

	fb "github.com/frabits/frabit-go-sdk/frabit"
	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) applyUser(group *gin.RouterGroup) {
	subRouter := group.Group("/users")
	subRouter.POST("", hs.CreateUser)
	subRouter.POST("/verify-email", hs.VerifyEmailStart)
	subRouter.GET("", hs.authZ(ac.EvalPermission(ac.ActionUserRead)), hs.GetUsers)
	subRouter.GET("/:login", hs.authZ(ac.EvalPermission(ac.ActionUserRead)), hs.GetUserByLogin)
	subRouter.DELETE("/:login", hs.DeleteUserByLogin)
}

func (hs *HttpServer) CreateUser(c *gin.Context) {
	var userInfo fb.CreateUserRequest
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		hs.Logger.Error("Bind team info failed", "Error", err.Error())
	}
	_, err := hs.user.CreateUser(hs.ctx, userInfo)
	if err != nil {
		hs.Logger.Error("create user failed", "username", userInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "create user failed")
	}
	auditCmd := &audit.CreateAuditCmd{
		Username:  "admin",
		EventName: event.EvtCreateUser,
		ClientIp:  c.ClientIP(),
	}
	hs.Logger.Info("write audit log", "eventName", auditCmd.EventName)
	if err := hs.audit.AddAuditEvent(hs.ctx, auditCmd); err != nil {
		hs.Logger.Error("add audit event failed")
	}
	c.IndentedJSON(http.StatusOK, "create an user successfully")
}

func (hs *HttpServer) VerifyEmailStart(c *gin.Context) {
	login := c.GetString(authn.AuthIdentity)
	if err := hs.verifier.Start(hs.ctx, login); err != nil {
		hs.Logger.Error("send verify request failed", "Error", err.Error())
		c.IndentedJSON(http.StatusOK, map[string]any{
			"title": "send verify request failed",
			"mesg":  err.Error(),
		})
		return
	}

	auditCmd := &audit.CreateAuditCmd{
		Username:  login,
		EventName: event.EvtCreateUser,
		ClientIp:  c.ClientIP(),
	}
	hs.Logger.Info("write audit log", "eventName", auditCmd.EventName)
	if err := hs.audit.AddAuditEvent(hs.ctx, auditCmd); err != nil {
		hs.Logger.Error("add audit event failed")
	}
	c.IndentedJSON(http.StatusOK, "create an user successfully")
}

func (hs *HttpServer) VerifyEmailComplete(c *gin.Context) {
	code := c.Param("code")
	if err := hs.verifier.Complete(hs.ctx, code); err != nil {
		hs.Logger.Error("send verify request failed", "Error", err.Error())
		c.IndentedJSON(http.StatusOK, map[string]any{
			"title": "verify email failed",
			"mesg":  err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, "email verified")
}

func (hs *HttpServer) GetUsers(c *gin.Context) {
	users, err := hs.user.GetUsers(hs.ctx)
	if err != nil {
		hs.Logger.Error("query user failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "query user failed")
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (hs *HttpServer) GetUserByLogin(c *gin.Context) {
	login := c.Param("login")
	user, err := hs.user.GetUserByLogin(hs.ctx, login)
	if err != nil {
		hs.Logger.Error("query user failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "query user failed")
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (hs *HttpServer) DeleteUserByLogin(c *gin.Context) {
	login := c.Param("login")
	err := hs.user.DeleteUser(hs.ctx, login)
	if err != nil {
		hs.Logger.Error("query user failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "query user failed")
	}
	c.IndentedJSON(http.StatusOK, "remove user successfully")
}
