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
	sa "github.com/frabits/frabit/pkg/services/serviceaccount"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hs *HttpServer) applyServiceAccount(group *gin.RouterGroup) {
	subRouter := group.Group("/sa")
	subRouter.POST("", hs.CreateServiceAccount)
	subRouter.GET("", hs.GetServiceAccount)
	subRouter.GET("/:name", hs.GetServiceAccountByName)
}

func (hs *HttpServer) GetServiceAccountByName(c *gin.Context) {
	login := c.Param("name")
	saAccounts, err := hs.serviceAccount.GetServiceAccountByName(hs.ctx, login)
	if err != nil {
		hs.Logger.Error("query service account failed", "Name", login)
		c.IndentedJSON(http.StatusBadGateway, "query service account failed")
	}
	c.IndentedJSON(http.StatusOK, saAccounts)
}

func (hs *HttpServer) GetServiceAccount(c *gin.Context) {
	login := c.Param("name")
	saAccounts, err := hs.serviceAccount.GetServiceAccount(hs.ctx)
	if err != nil {
		hs.Logger.Error("query service account failed", "Name", login)
		c.IndentedJSON(http.StatusBadGateway, "query service account failed")
	}
	c.IndentedJSON(http.StatusOK, saAccounts)
}

func (hs *HttpServer) CreateServiceAccount(c *gin.Context) {
	var saInfo sa.CreateServiceAccountCmd
	if err := c.ShouldBindJSON(&saInfo); err != nil {
		hs.Logger.Error("Bind service account info failed", "Error", err.Error())
	}
	if err := hs.serviceAccount.CreateServiceAccount(hs.ctx, &saInfo); err != nil {
		hs.Logger.Error("create service account failed", "Name", saInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "create service account failed")
	}
	c.IndentedJSON(http.StatusOK, "create service account successfully")
}
