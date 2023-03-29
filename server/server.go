// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
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

package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/frabits/frabit/common/log"
	"github.com/frabits/frabit/server/config"
	store "github.com/frabits/frabit/server/meta_store"
	"github.com/frabits/frabit/server/router"
	"github.com/frabits/frabit/server/service"
	"github.com/frabits/frabit/server/service/backup"
	"github.com/frabits/frabit/server/service/deploy"
)

type Server struct {
	startedTs      int64
	BackupService  backup.BackupService
	RestoreService service.RestoreService
	DeployService  deploy.DeployService
	UpgradeService service.UpgradeService
	config         config.Config
	g              router.Router
	db             *sql.DB
}

func NewServer(cfg config.Config) *Server {
	meta, err := store.OpenFrabit()
	if err != nil {

	}
	srv := &Server{
		startedTs: time.Now().Unix(),
		config:    cfg,
		db:        meta,
	}

	return srv
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.g.Run(s.config.Server.Port); err != nil {
		log.Error("Server start failed")
		return err

	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.g.Run(s.config.Server.Port); err != nil {
		log.Error("Server start failed")
		return err

	}
	return nil
}

func (s *Server) initSubscription() {
	s.getLicense()
}

func (s *Server) getLicense() {
	fmt.Println("unImplement")
}
