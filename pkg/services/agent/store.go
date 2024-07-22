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
	"github.com/frabits/frabit/pkg/infra/log"
	"gorm.io/gorm"
	"log/slog"
)

type store interface {
	CreateAgent(context.Context, *Agent) (uint32, error)
	GetAgents(context.Context) ([]Agent, error)
	GetAgentById(context.Context, string) (Agent, error)
	RemoveAgentById(context.Context, string) error
	SendHeartbeat(context.Context, *Heartbeat) error
}

type storeImpl struct {
	DB  *gorm.DB
	log *slog.Logger
}

func newStoreImpl(db *gorm.DB) *storeImpl {
	return &storeImpl{
		DB:  db,
		log: log.New("agent.store"),
	}
}

func (s *storeImpl) CreateAgent(ctx context.Context, agent *Agent) (uint32, error) {
	s.DB.Create(agent)
	return 0, nil
}

func (s *storeImpl) GetAgents(ctx context.Context) ([]Agent, error) {
	var agents []Agent
	s.DB.Model(Agent{}).Find(&agents)
	return agents, nil
}

func (s *storeImpl) GetAgentById(ctx context.Context, agentId string) (Agent, error) {
	var agent Agent
	s.DB.Model(Agent{}).Where("agent_id = ?", agentId).First(&agent)
	return agent, nil
}

func (s *storeImpl) RemoveAgentById(ctx context.Context, agentId string) error {
	var agent Agent
	s.DB.Where("agent_id = ?", agentId).Delete(&agent)
	return nil
}

func (s *storeImpl) SendHeartbeat(ctx context.Context, beat *Heartbeat) error {
	var hb []Heartbeat
	s.DB.Model(Heartbeat{}).Where("agent_id = ?", beat.AgentId).Find(&hb)
	if len(hb) > 0 {
		s.DB.Model(Heartbeat{}).Where("agent_id = ?", beat.AgentId).Updates(Heartbeat{Status: beat.Status, UpdatedAt: beat.UpdatedAt})
	} else {
		s.DB.Create(beat)
	}
	return nil
}
