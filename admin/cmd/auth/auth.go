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

package auth

import (
	"github.com/spf13/cobra"
)

// CmdAuth authenticate frabit-admin and frabit-server via token
var CmdAuth = &cobra.Command{
	Use:   "auth <subcommand> [flag]",
	Short: "frabit auth manager",
}

func init() {

	CmdAuth.AddCommand(CmdLogin)
	CmdAuth.AddCommand(CmdLogout)
	CmdAuth.AddCommand(CmdStatus)
}
