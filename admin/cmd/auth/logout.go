// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Blylei <blylei.info@gmail.com>
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
	"fmt"
	"github.com/frabits/frabit/common/version"
	"github.com/spf13/cobra"
)

var CmdLogout = &cobra.Command{
	Use:   "logout [flag]",
	Short: "Logout frabit-admin from frabit-server,remove token from keyring/env ",
	Run:   runLogout,
}

func runLogout(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
}

func init() {
	// cmd.rootCmd.AddCommand(newVersionCmd)
}
