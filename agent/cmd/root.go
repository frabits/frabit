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

package cmd

import (
	"fmt"
	"os"

	"github.com/frabits/frabit/pkg/common/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "frabit-agent",
	Short:   "A component used with frabit-server",
	Run:     runAgent,
	Version: version.InfoStr.String(),
}

var flag struct {
	port int
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
	rootCmd.PersistentFlags().IntVar(&flag.port, "port", 80, "port. Default to 9180")
}

func runAgent(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
}
