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

	"github.com/gin-gonic/gin"

	fb "github.com/frabits/frabit-go-sdk/frabit"
)

func (hs *HttpServer) applyAgent(group *gin.RouterGroup) {
	agentRouter := group.Group("/agent")
	agentRouter.POST("/register", hs.Register)
	agentRouter.POST("/remove", hs.UnRegister)
	agentRouter.GET("/list", hs.GetAgents)
	agentRouter.GET("/list/:agent_id", hs.GetAgents)
	agentRouter.POST("/heartbeat", hs.Heartbeat)
}

func (hs *HttpServer) Register(c *gin.Context) {
	var agentInfo fb.CreateAgentRequest
	if err := c.ShouldBindJSON(&agentInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}
	if err := hs.agent.Register(hs.ctx, agentInfo); err != nil {
		hs.Logger.Error("register agent failed", "agent_id", agentInfo.AgentID)
		c.IndentedJSON(http.StatusBadGateway, "register agent failed")
	}
	c.IndentedJSON(http.StatusOK, "register an agent")
}

func (hs *HttpServer) UnRegister(c *gin.Context) {
	agentId := c.Query("agent_id")
	err := hs.agent.UnRegister(hs.ctx, agentId)
	if err != nil {
		hs.Logger.Error("remove agent failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "remove agent failed")
	}
	c.IndentedJSON(http.StatusOK, "agent removed successfully")
}

func (hs *HttpServer) GetAgents(c *gin.Context) {
	agents, err := hs.agent.GetAgents(hs.ctx)
	if err != nil {
		hs.Logger.Error("query agent list failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "get agent list failed")
	}
	c.IndentedJSON(http.StatusOK, agents)
}

func (hs *HttpServer) GetAgentById(c *gin.Context) {
	agentId := c.Param("agent_id")
	agent, err := hs.agent.GetAgentById(hs.ctx, agentId)
	if err != nil {
		hs.Logger.Error("query agent list failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "get agent list failed")
	}
	c.IndentedJSON(http.StatusOK, agent)
}

func (hs *HttpServer) Heartbeat(c *gin.Context) {
	var beatInfo fb.CreateHeartbeat
	if err := c.ShouldBindJSON(&beatInfo); err != nil {
		hs.Logger.Error("Bind heartbeat info failed", "Error", err.Error())
	}
	hs.agent.Heartbeat(hs.ctx, beatInfo)
	c.IndentedJSON(http.StatusOK, "send heartbeat successfully")
}
