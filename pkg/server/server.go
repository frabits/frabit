// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
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
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/frabits/frabit/pkg/api"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"
	"github.com/frabits/frabit/pkg/registry"
	"github.com/frabits/frabit/pkg/server/bg_services"
)

type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	childRoutines    *errgroup.Group
	log              *slog.Logger
	mtx              sync.Mutex
	startedTs        int64
	License          string

	pidFile string
	version string

	backgroundServices []registry.BackgroundService

	config     *config.Config
	httpServer *api.HttpServer
	db         *sql.DB
}

func (s *Server) init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.initSubscription()
	s.writePIDFile()

	return nil
}

func NewServer(cfg *config.Config) *Server {
	backgroundService := bg_services.BgSvcs

	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)
	httpServer := api.NewHttpServer(cfg)
	metaStore, err := db.New(cfg)
	if err != nil {
		fmt.Println("Create metastore failed", err.Error())
		os.Exit(1)
	}
	stdDB, err := metaStore.OpenFrabit()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	srv := &Server{
		startedTs:          time.Now().Unix(),
		context:            childCtx,
		childRoutines:      childRoutines,
		shutdownFn:         shutdownFn,
		shutdownFinished:   make(chan struct{}),
		backgroundServices: backgroundService.GetServices(),
		config:             cfg,
		log:                log.New("Server"),

		httpServer: httpServer,
		db:         stdDB,
	}

	return srv
}

func (s *Server) Run() error {
	defer close(s.shutdownFinished)
	if err := s.init(); err != nil {
		return err
	}

	servicesHub := s.backgroundServices
	// Start each background service.
	for _, svc := range servicesHub {
		service := svc

		serviceName := reflect.TypeOf(service).String()
		s.childRoutines.Go(func() error {
			select {
			case <-s.context.Done():
				return s.context.Err()
			default:
			}
			s.log.Debug("Starting background service", "service", serviceName)
			err := service.Run(s.context)
			// Do not return context.Canceled error
			if err != nil && errors.Is(err, context.Canceled) {
				s.log.Error("Stopped background service", "Error", err.Error())
				return fmt.Errorf("%s run error:%w", serviceName, err)
			}
			s.log.Debug("Stopped background service", "service", serviceName)
			return nil
		})
	}
	// finally, start http server
	s.childRoutines.Go(func() error {
		select {
		case <-s.context.Done():
			return s.context.Err()
		default:
		}
		s.log.Debug("Starting background service", "service", "httpserver")
		err := s.httpServer.Run(s.context)
		// Do not return context.Canceled error
		if err != nil && errors.Is(err, context.Canceled) {
			s.log.Error("Stopped background service", "Error", err.Error())
			return fmt.Errorf("%s run error:%w", "httpserver", err)
		}
		s.log.Debug("Stopped background service", "service", "httpserver")
		return nil
	})
	s.notifySystemd("READY=1")
	s.log.Info("Waiting on services...", "Port", s.config.Server.Port)
	fmt.Printf("Server started at %v:%v\n", "http://localhost", s.config.Server.Port)
	return s.childRoutines.Wait()
}

// Shutdown initializes  Frabit graceful shutdown. This shuts down all running background services.
// Since Run blocks supposed to be run from a separate goroutine
func (s *Server) Shutdown(ctx context.Context, reason string) error {
	var err error
	// firstly, shutdown httpServer to avoid new http request
	if err = s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	s.shutdownOnce.Do(func() {
		s.log.Info("Shutdown started", "reason", reason)
		// revoke cancel function to stop services
		s.shutdownFn()

		// Wait for server to shut down
		select {
		case <-s.shutdownFinished:
			s.log.Debug("Finished waiting for server to shut down")
		case <-ctx.Done():
			s.log.Warn("Time out while waiting for server to shut down")
			err = fmt.Errorf("timeout waiting for shutdown")
		}
	})
	return err
}

func (s *Server) initSubscription() {
	s.getLicense()
}

func (s *Server) getLicense() {
	s.License = "Ultimate"
}

func (s *Server) writePIDFile() error {
	if s.pidFile == "" {
		s.pidFile = "/var/run/frabit-server.pid"
	}

	err := os.MkdirAll(filepath.Dir(s.pidFile), 0700)
	if err != nil {
		s.log.Error("Failed to verify pid directory")
		return fmt.Errorf("failed to verify pid directory:%s", err)
	}

	// Get the pid and write it to file
	pid := strconv.Itoa(os.Getpid())

	if err := os.WriteFile(s.pidFile, []byte(pid), 0644); err != nil {
		s.log.Error("Failed to write pidfile")
		return fmt.Errorf("failed to write pidfile:%s", err)
	}
	s.log.Info("Write PID file")

	return nil
}

func (s *Server) notifySystemd(state string) {

}
