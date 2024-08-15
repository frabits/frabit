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
	"github.com/frabits/frabit/pkg/services/apikey"
	"github.com/frabits/frabit/pkg/services/license"
	"github.com/frabits/frabit/pkg/services/settings"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/frabits/frabit/pkg/bus"
	"github.com/frabits/frabit/pkg/common/version"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/services/agent"
	"github.com/frabits/frabit/pkg/services/audit"
	"github.com/frabits/frabit/pkg/services/auth"
	"github.com/frabits/frabit/pkg/services/backup"
	"github.com/frabits/frabit/pkg/services/deploy"
	"github.com/frabits/frabit/pkg/services/login"
	"github.com/frabits/frabit/pkg/services/org"
	sa "github.com/frabits/frabit/pkg/services/serviceaccount"
	"github.com/frabits/frabit/pkg/services/team"
	"github.com/frabits/frabit/pkg/services/user"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	ctx    context.Context
	Logger *slog.Logger
	Server *http.Server
	router *gin.Engine
	Port   uint32
	Bus    bus.Bus

	apiKey         apikey.Service
	authUser       auth.Service
	audit          audit.Service
	backup         backup.Service
	deploy         deploy.Service
	agent          agent.Service
	login          login.Service
	license        license.Service
	org            org.Service
	team           team.Service
	user           user.Service
	serviceAccount sa.Service
	settings       settings.Service
}

func ProviderHTTPServer(cnf *config.Config, auth auth.Service, backup backup.Service, deploy deploy.Service, agent agent.Service,
	login login.Service, org org.Service, team team.Service, user user.Service, audit audit.Service, bus bus.Bus, sa sa.Service,
	apikey apikey.Service, settings settings.Service, license license.Service) *HttpServer {
	var port uint32

	if cnf.Server.Port != 0 {
		port = cnf.Server.Port
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	logfile, _ := os.OpenFile(cnf.Logger.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	router.Use(gin.LoggerWithWriter(logfile))
	srv := &http.Server{
		Addr:    ":9180",
		Handler: router.Handler(),
	}

	hs := &HttpServer{
		Logger: log.New("http.server"),
		Server: srv,
		Port:   port,
		Bus:    bus,
		router: router,

		apiKey:         apikey,
		authUser:       auth,
		audit:          audit,
		backup:         backup,
		deploy:         deploy,
		license:        license,
		org:            org,
		team:           team,
		user:           user,
		agent:          agent,
		login:          login,
		serviceAccount: sa,
		settings:       settings,
	}

	return hs
}

// setup register all router
func (hs *HttpServer) setup(engine *gin.Engine) {
	engine.GET("/health", hs.health)
	engine.GET("/info", hs.info)
	engine.POST("/login", hs.Login)
	apiV2 := engine.Group("/api/v2")
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

func (hs *HttpServer) health(c *gin.Context) {
	info := struct {
		Timestamp time.Time `json:"timestamp"`
		Database  string    `json:"database"`
		APIServer string    `json:"api_server"`
	}{Timestamp: time.Now(),
		Database:  "ok",
		APIServer: "ok",
	}
	hs.Logger.Info("Hearth check revoked")
	c.IndentedJSON(http.StatusOK, info)
}

func (hs *HttpServer) info(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, version.InfoStr)
}

func (hs *HttpServer) Run(ctx context.Context) error {
	hs.Logger.Info("Start http.server")
	hs.setup(hs.router)
	hs.Logger.Info("setup api endpoints")
	return hs.Server.ListenAndServe()
}

func (hs *HttpServer) Shutdown(ctx context.Context) error {
	hs.Logger.Info("start shutdown httpServer")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	cancel()
	return hs.Server.Shutdown(ctx)
}
