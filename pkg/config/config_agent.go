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

package config

import "log/slog"

type AgentConfig struct {
	ServerURL    string
	Port         uint32
	FileName     string
	Format       string
	Security     string
	DefaultLevel slog.Level
	MaxDay       uint
}

var AgentConf = newAgentConfig()

func newAgentConfig() *AgentConfig {
	agentConf := &AgentConfig{
		ServerURL:    "http://localhost:9180",
		Port:         19180,
		FileName:     "/tmp/frabit_agent.log",
		Format:       "json",
		DefaultLevel: slog.LevelInfo,
	}
	return agentConf
}
