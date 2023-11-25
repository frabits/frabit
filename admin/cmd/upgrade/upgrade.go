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

package upgrade

import (
	"github.com/frabits/frabit/pkg/common/cmdutil"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the backup command
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade <command>",
	Short: "Upgrade manager",
}

func init() {
	cmdutil.AddGroup(CmdUpgrade, "Upgrade commands", mysqlCmd, clickhouseCmd, redisCmd, mongodbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CmdBackup.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CmdBackup.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
