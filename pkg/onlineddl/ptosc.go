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

package onlineddl

import (
	"github.com/frabits/frabit/pkg/common/constant"
	"github.com/frabits/frabit/pkg/common/log"
	"os/exec"
	"time"

	"go.uber.org/zap"
)

// PTOSC represent a pt-online-schema-change provided by Percona LLC
type PTOSC struct {
	cmd       string
	StartTime time.Time
	opt       ptoscOpt
	logger    *zap.Logger
}

type ptoscOpt struct {
	user     string
	password string
	host     string
	port     uint32

	alterStatement string
}

func NewPTOSC() *PTOSC {
	cmdPT, err := exec.LookPath(constant.PTOSC)
	if err != nil {
		log.Info("pt-osc not in you path environment")
	}
	return &PTOSC{
		cmd:    cmdPT,
		logger: log.Logger,
	}
}

func init() {
	NewPTOSC()
}

func (pt *PTOSC) Run() error {
	pt.StartTime = time.Now()
	cmd := exec.Cmd{Path: pt.cmd}
	return cmd.Run()
}
