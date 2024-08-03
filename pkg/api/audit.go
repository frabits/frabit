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
)

func (hs *HttpServer) applyAudit(group *gin.RouterGroup) {
	subRouter := group.Group("/audits")
	subRouter.GET("/list", hs.GetAudit)
}

func (hs *HttpServer) GetAudit(c *gin.Context) {
	audits, err := hs.audit.GetAuditEvent(hs.ctx)
	if err != nil {
		hs.Logger.Error("query agent list failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "get audit list failed")
	}
	c.IndentedJSON(http.StatusOK, audits)
}
