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
	"log/slog"
	"time"

	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/db"
	"github.com/frabits/frabit/pkg/infra/log"

	fb "github.com/frabits/frabit-go-sdk/frabit"
)

type agentService struct {
	cfg   *config.Config
	store store
	log   *slog.Logger
}

func ProviderAgentService(cfg *config.Config, metaDB *db.MetaStore) Service {
	metaStore := newStoreImpl(metaDB.Gdb)
	as := &agentService{
		cfg:   cfg,
		store: metaStore,
		log:   log.New("agent"),
	}

	return as
}

func (s *agentService) Register(ctx context.Context, agent fb.CreateAgentRequest) error {
	datetime := time.Now().Format(time.DateTime)
	agentInstance := &Agent{
		Hostname:  agent.Name,
		AgentId:   agent.AgentID,
		ClientIp:  agent.ClientIP,
		CreatedAt: datetime,
		UpdatedAt: datetime,
	}
	aid, err := s.store.CreateAgent(ctx, agentInstance)
	if err != nil {
		s.log.Error("create agent failed", "Error", err.Error())
		return err
	}
	s.log.Info("create agent successfully", "id", aid)
	return nil
}

func (s *agentService) UnRegister(ctx context.Context, agentId string) error {
	return s.store.RemoveAgentById(ctx, agentId)
}

func (s *agentService) Heartbeat(ctx context.Context, heartbeat fb.CreateHeartbeat) error {
	datetime := time.Now().Format(time.DateTime)
	hb := &Heartbeat{
		AgentId:   heartbeat.AgentID,
		Status:    heartbeat.Status,
		CreatedAt: datetime,
		UpdatedAt: datetime,
	}
	if err := s.store.SendHeartbeat(ctx, hb); err != nil {
		s.log.Debug("send heartbeat failed", "Error", err.Error())
	}
	return nil
}

func (s *agentService) GetAgents(ctx context.Context) ([]Agent, error) {
	agents, err := s.store.GetAgents(ctx)
	if err != nil {
		s.log.Error("get agent list failed", "Error", err.Error())
		return agents, err
	}
	s.log.Debug("query agent successfully")
	return agents, nil
}

func (s *agentService) GetAgentById(ctx context.Context, agentId string) (Agent, error) {
	agent, err := s.store.GetAgentById(ctx, agentId)
	if err != nil {
		s.log.Error("get agent failed", "Error", err.Error())
		return agent, err
	}
	s.log.Debug("query agent successfully")
	return agent, nil
}
