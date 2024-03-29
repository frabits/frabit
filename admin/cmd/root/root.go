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

package root

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/frabits/frabit/admin/cmd/admin"
	"github.com/frabits/frabit/admin/cmd/auth"
	"github.com/frabits/frabit/admin/cmd/backup"
	"github.com/frabits/frabit/admin/cmd/deploy"
	"github.com/frabits/frabit/admin/cmd/plugin"
	"github.com/frabits/frabit/admin/cmd/restore"
	"github.com/frabits/frabit/admin/cmd/upgrade"
	"github.com/frabits/frabit/admin/cmd/version"
	"github.com/frabits/frabit/pkg/client"
	"github.com/frabits/frabit/pkg/common/cmdutil"
	"github.com/frabits/frabit/pkg/common/config"
)

type rootOpt struct {
	help bool
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "frabit-admin [subcommand] -flag",
	Short: "A CLI application to directly send API request to frabit-server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmdutil.IsAuthCheckEnabled(cmd) && !cmdutil.CheckAuth(config.Config{}) {
			fmt.Fprint(os.Stderr, authHelp())
			return errors.New("authError")
		}
		return nil
	},
}

func authHelp() string {
	return "frabit: To use frabit-admin CLI in automation, set the FB_TOKEN environment variable"
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(backup.CmdBackup)
	rootCmd.AddCommand(restore.CmdRestore)
	rootCmd.AddCommand(deploy.CmdDeploy)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(admin.ResetAdminPassword)
	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(plugin.CmdPlugin)
	rootCmd.AddCommand(auth.CmdAuth)
	_ = client.NewClient("localhost:9180")
	// rootCmd.PersistentFlags().BoolVar(&rootOpt.help, "help", false, "Show help information for each command")

}
