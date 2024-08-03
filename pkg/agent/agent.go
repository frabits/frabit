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

package agent

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/frabits/frabit/pkg/agent/executor"
	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"

	fb "github.com/frabits/frabit-go-sdk/frabit"
)

type Agent struct {
	agentId  string
	uuidFile string
	service  fb.AgentService
	Executor executor.Executor
	cfg      *config.AgentConfig
	Log      *slog.Logger
}

func New(cfg *config.AgentConfig) *Agent {
	if cfg == nil {
		cfg = config.AgentConf
	}
	client, err := fb.NewClient(fb.WithBaseURL(cfg.Server.ServerURL), fb.WithUserAgent("Agent"))
	if err != nil {
		fmt.Println("agent service start failed", "Error", err.Error())
	}
	return &Agent{
		cfg:     cfg,
		Log:     log.New("agent"),
		service: client.Agent,
	}
}

func (a *Agent) RunAgent(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return a.Run(ctx)
		}
	}
}

// Run initial
func (a *Agent) Run(ctx context.Context) error {
	a.Log.Info("start agent", "server", a.cfg.Server.ServerURL)
	// if agentId file not exist,we need register is firstly
	if _, err := os.Stat(a.uuidFile); os.IsNotExist(err) {
		a.Log.Info("agent not registered,auto-register")
		if err := a.register(ctx); err != nil {
			a.Log.Error("register agent failed", "Error", err.Error())
			return err
		}
	}
	if err := a.readAgentID(); err != nil {
		a.Log.Error("read agent id failed", "Error", err.Error())
		return err
	}

	// health check via send heartbeat to server
	go func() {
		tick := time.NewTicker(10 * time.Second)
		defer tick.Stop()
		for {
			select {
			case <-ctx.Done():
				a.Log.Info("shutdown service successfully")
			case <-tick.C:
				a.heartbeat(ctx)
			}
		}
	}()

	return nil
}

func (a *Agent) Shutdown(ctx context.Context, reason string) error {
	a.Log.Info("shutdown agent", "server", a.cfg.Server.ServerURL)
	return nil
}

func (a *Agent) register(ctx context.Context) error {
	agentId := utils.CreateUUIDWithDelimiter("")
	agent := fb.CreateAgentRequest{
		AgentID: agentId,
		Name:    "dbaDest" + time.Now().Format("200601021504"),
		Status:  "",
	}
	a.Log.Info("agent info", "Agent", agent)
	if err := a.service.Register(ctx, agent); err != nil {
		a.Log.Error("register agent failed via sdk", "Error", err.Error())
		return err
	}
	a.agentId = agentId
	a.Log.Info("add agent", "name", agent.Name, "agent_id", agent.AgentID)
	if err := a.writeAgentID(); err != nil {
		a.Log.Error("persist agentId failed", "Error", err.Error())
		return err
	}
	return nil
}

func (a *Agent) heartbeat(ctx context.Context) error {
	heartbeat := fb.CreateHeartbeat{
		AgentID: a.agentId,
		Status:  "ok",
	}
	a.Log.Info("send heartbeat to server", "agent_id", heartbeat.AgentID)
	return a.service.Heartbeat(ctx, heartbeat)
}

func (a *Agent) writeAgentID() error {
	a.uuidFile = "/usr/local/frabit/frabit-agent.txt"
	err := os.MkdirAll(filepath.Dir(a.uuidFile), 0700)
	if err != nil {
		a.Log.Error("Failed to verify pid directory")
		return fmt.Errorf("failed to verify pid directory:%s", err)
	}

	// Get the agentId and write it to file
	agentId := a.agentId

	if err := os.WriteFile(a.uuidFile, []byte(agentId), 0644); err != nil {
		a.Log.Error("Failed to write pid file", "Error", err.Error())
		return fmt.Errorf("failed to write pidfile:%s", err)
	}
	a.Log.Info("Write uuid file")

	return nil
}

func (a *Agent) readAgentID() error {
	a.uuidFile = "/usr/local/frabit/agent.txt"
	agentId, err := os.ReadFile(a.uuidFile)
	if err != nil {
		a.Log.Error("read agent_id failed", "Error", err.Error())
		return err
	}
	a.agentId = string(agentId)

	return nil
}
