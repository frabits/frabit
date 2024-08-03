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

func (hs *HttpServer) applyTeam(group *gin.RouterGroup) {
	subRouter := group.Group("/teams")
	subRouter.POST("", hs.CreateTeam)
	subRouter.GET("", hs.GetTeams)
	subRouter.GET("/:name", hs.GetTeamByName)
}

func (hs *HttpServer) CreateTeam(c *gin.Context) {
	var teamInfo fb.CreateTeamRequest
	if err := c.ShouldBindJSON(&teamInfo); err != nil {
		hs.Logger.Error("Bind team info failed", "Error", err.Error())
	}
	if err := hs.team.Create(hs.ctx, teamInfo); err != nil {
		hs.Logger.Error("register agent failed", "team_name", teamInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "create team failed")
	}
	c.IndentedJSON(http.StatusOK, "create an team successfully")
}

func (hs *HttpServer) GetTeams(c *gin.Context) {
	teams, err := hs.team.GetAll(hs.ctx)
	if err != nil {
		hs.Logger.Error("query team failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "query team failed")
	}
	c.IndentedJSON(http.StatusOK, teams)
}

func (hs *HttpServer) GetTeamByName(c *gin.Context) {
	name := c.Param("name")
	team, err := hs.team.GetTeamByName(hs.ctx, name)
	if err != nil {
		hs.Logger.Error("query team failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "query team failed")
	}
	c.IndentedJSON(http.StatusOK, team)
}
