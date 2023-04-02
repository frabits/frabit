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

package deploy

import (
	"fmt"
	"github.com/frabits/frabit/common/version"
	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Deploy a mysql database based on provide topology",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	mysqlCmd.AddCommand(cmdStandalone)
	mysqlCmd.AddCommand(cmdReplicate)
}

var cmdStandalone = &cobra.Command{
	Use:   "standalone",
	Short: "Deploy a standalone mysql instance",
	Run:   runStandalone,
}

var cmdReplicate = &cobra.Command{
	Use:   "replication",
	Short: "Deploy a mysql replication topology",
	Run:   runReplicate,
}

func runStandalone(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
}

func runReplicate(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
}
