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
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/registry"
	"github.com/frabits/frabit/pkg/server/bg_services"
	"github.com/frabits/frabit/pkg/server/router"
)

type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	childRoutines    *errgroup.Group
	log              *zap.Logger
	mtx              sync.Mutex
	startedTs        int64

	pidFile string
	version string

	backgroundServices []registry.BackgroundService

	config config.Config
	g      router.Router
	db     *sql.DB
}

func (s *Server) init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.initSubscription()

	return nil
}

func NewServer(cfg config.Config) *Server {
	//meta, err := store.OpenFrabit()
	//if err != nil {
	//
	//}
	backgroundService := bg_services.BgSvcs

	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)
	srv := &Server{
		startedTs:          time.Now().Unix(),
		context:            childCtx,
		childRoutines:      childRoutines,
		shutdownFn:         shutdownFn,
		backgroundServices: backgroundService.GetServices(),
		config:             cfg,
		//db:                 meta,
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
			s.log.Debug("Starting background service", zap.String("service", serviceName))
			err := service.Run(s.context)
			// Do not return context.Canceled error
			if err != nil && errors.Is(err, context.Canceled) {
				s.log.Error("Stopped background service", zap.Error(err))
				return fmt.Errorf("%s run error:%w", serviceName, err)
			}
			s.log.Debug("Stopped background service", zap.String("service", serviceName))
			return nil
		})
	}
	s.notifySystemd("READY=1")
	s.log.Debug("Waiting on services...")
	return s.childRoutines.Wait()
}

// Shutdown initializes  Frabit graceful shutdown. This shuts down all running background services.
// Since Run blocks supposed to be run from a separate goroutine
func (s *Server) Shutdown(ctx context.Context, reason string) error {
	var err error
	s.shutdownOnce.Do(func() {
		s.log.Info("Shutdown started", zap.String("reason", reason))
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
	fmt.Println("getLicense unImplement")
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
