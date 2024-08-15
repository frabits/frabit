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
	"github.com/frabits/frabit/pkg/satoken"
	"net/http"

	"github.com/frabits/frabit/pkg/services/apikey"

	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) applyApiKey(group *gin.RouterGroup) {
	subRouter := group.Group("/apikey")
	subRouter.POST("", hs.CreateApiKey)
}

func (hs *HttpServer) CreateApiKey(c *gin.Context) {
	var keyInfo apikey.CreateAPIKeyCmd
	if err := c.ShouldBindJSON(&keyInfo); err != nil {
		hs.Logger.Error("Bind agent info failed", "Error", err.Error())
	}

	newKey := satoken.New()
	keyInfo.HashKey = newKey.HashedKey
	metaKey, err := hs.apiKey.AddAPIKey(hs.ctx, &keyInfo)
	if err != nil {
		hs.Logger.Error("create apikey failed", "Error", err.Error())
		c.IndentedJSON(http.StatusBadGateway, "create apikey failed")
	}
	result := apikey.APIKeyDTO{
		Name:         metaKey.Name,
		ClientSecret: newKey.ClientSecret,
	}
	c.IndentedJSON(http.StatusOK, result)
}
