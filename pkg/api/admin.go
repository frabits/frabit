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

	ac "github.com/frabits/frabit/pkg/services/access_control"
	"github.com/frabits/frabit/pkg/services/settings"
	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) applyAdmin(group *gin.RouterGroup) {
	subRouter := group.Group("/admin")
	subRouter.POST("/settings", hs.authZ(ac.EvalPermission(ac.ActionSettingWrite)), hs.CreateSetting)
	subRouter.GET("/settings", hs.authZ(ac.EvalPermission(ac.ActionSettingRead)), hs.GetAllSetting)
}

func (hs *HttpServer) CreateSetting(c *gin.Context) {
	var settingsInfo settings.CreateSettingsCmd
	if err := c.ShouldBindJSON(&settingsInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}
	if err := hs.settings.CreateSettings(hs.ctx, &settingsInfo); err != nil {
		hs.Logger.Error("create sso setting failed", "Name", settingsInfo.Name)
		c.IndentedJSON(http.StatusBadGateway, "register agent failed")
	}
	c.IndentedJSON(http.StatusOK, "register an agent")
}

func (hs *HttpServer) GetAllSetting(c *gin.Context) {
	results, err := hs.settings.GetSettings(hs.ctx)
	if err != nil {
		hs.Logger.Error("create sso setting failed")
		c.IndentedJSON(http.StatusBadGateway, "register agent failed")
	}
	c.IndentedJSON(http.StatusOK, results)
}
