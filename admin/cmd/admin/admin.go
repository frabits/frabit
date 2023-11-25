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

package admin

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdAdmin represents the Admin command
var CmdAdmin = &cobra.Command{
	Use:   "reset-admin-password <new_password>",
	Short: "Reset admin password at any time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("standalone called")
	},
}

var resetAdminPassword = &cobra.Command{
	Use:   "reset-admin-password <new_password>",
	Short: "Reset admin password at any time",
	Run:   runResetAdminPassword,
}

func runResetAdminPassword(cmd *cobra.Command, args []string) {
	fmt.Println("runResetAdminPassword called")
}

func init() {
	CmdAdmin.AddCommand(resetAdminPassword)
	CmdAdmin.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
