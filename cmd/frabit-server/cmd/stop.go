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
	"context"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/server"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop frabit-server daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop() {
	ctx := context.Background()
	cfg := config.Conf
	srv := server.NewServer(cfg)
	srv.Shutdown(ctx, "")
}
