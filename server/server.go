/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package server

import (
	"context"
	"fmt"
	"time"

	"github.com/frabits/frabit/server/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	startedTs int64
	Conf      config.Config
	g         gin.Engine
}

func NewServer(ctx context.Context, cfg config.Config) *Server {
	srv := &Server{
		startedTs: time.Now().Unix(),
		Conf:      cfg,
	}

	return srv
}

func (s *Server) Run(ctx context.Context, port int) error {
	fmt.Println("unImplement")
	return nil
}

func (s *Server) Shutdown(ctx context.Context, port int) error {
	fmt.Println("unImplement")
	return nil
}

func (s *Server) initSubscription() {
	s.getLicense()
}

func (s *Server) getLicense() {
	fmt.Println("unImplement")
}
