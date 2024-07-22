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

	fb "github.com/frabits/frabit-go-sdk/frabit"
	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) applyOrg(group *gin.RouterGroup) {
	agentRouter := group.Group("/orgs")
	agentRouter.GET("", hs.GetOrgs)
	agentRouter.POST("", hs.CreateOrg)
	agentRouter.PUT("", hs.UpdateOrg)
}

func (hs *HttpServer) GetOrgs(c *gin.Context) {
	orgs, err := hs.org.GetOrgs(hs.ctx)
	if err != nil {
		hs.Logger.Error("query org list failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "get org list failed")
	}
	c.IndentedJSON(http.StatusOK, orgs)
}

func (hs *HttpServer) CreateOrg(c *gin.Context) {
	var orgInfo fb.OrgCreateRequest
	if err := c.ShouldBindJSON(&orgInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}

	_, err := hs.org.CreateOrg(hs.ctx, orgInfo)
	if err != nil {
		hs.Logger.Error("create org failed", "org_name", orgInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "create org failed")
	}
	c.IndentedJSON(http.StatusOK, "create an org successfully")
}

func (hs *HttpServer) UpdateOrg(c *gin.Context) {
	var orgInfo fb.OrgUpdateRequest
	if err := c.ShouldBindJSON(&orgInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}

	err := hs.org.UpdateOrg(hs.ctx, orgInfo)
	if err != nil {
		hs.Logger.Error("create org failed", "org_name", orgInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "create org failed")
	}
	c.IndentedJSON(http.StatusOK, "register an agent")
}
