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
	"context"
	"github.com/frabits/frabit/pkg/infra/log"
	"log/slog"
	"net/http"
	"time"

	"github.com/frabits/frabit/pkg/common/version"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/services/backup"
	"github.com/frabits/frabit/pkg/services/deploy"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Logger *slog.Logger
	Server *http.Server
	Port   uint32

	backup backup.Service
	deploy deploy.Service
	// login  login.Service
}

func NewHttpServer(cnf *config.Config) *HttpServer {
	var port uint32

	if cnf.Server.Port != 0 {
		port = cnf.Server.Port
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	Setup(router)
	srv := &http.Server{
		Addr:    ":9180",
		Handler: router.Handler(),
	}

	hs := &HttpServer{
		Logger: log.New("http.server"),
		Server: srv,
		Port:   port,
	}

	return hs
}

// Setup register all router
func Setup(engine *gin.Engine) {
	engine.GET("/health", health)
	engine.GET("/info", info)
	engine.POST("/login", info)
	apiV2 := engine.Group("/api/v2")
	applyBackup(apiV2)
	applyDeploy(apiV2)
}

func health(c *gin.Context) {
	info := struct {
		Timestamp time.Time `json:"timestamp"`
		Database  string    `json:"database"`
		APIServer string    `json:"api_server"`
	}{Timestamp: time.Now(),
		Database:  "ok",
		APIServer: "ok",
	}
	c.IndentedJSON(http.StatusOK, info)
}

func info(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, version.InfoStr)
}

func (hs *HttpServer) Run(ctx context.Context) error {
	hs.Logger.Info("Start httpServer")
	return hs.Server.ListenAndServe()
}

func (hs *HttpServer) Shutdown(ctx context.Context) error {
	hs.Logger.Info("start shutdown httpServer")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	cancel()
	return hs.Server.Shutdown(ctx)
	/*
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer func() {
				hs.Logger.Info("shutdown http_server timeout")
				cancel()
			}()
			return hs.Server.Shutdown(ctx)
		default:
			return ctx.Err()
		}

	*/
}
