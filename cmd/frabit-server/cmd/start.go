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

package cmd

import (
	"context"

	"github.com/frabits/frabit/server"
	"github.com/frabits/frabit/server/config"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start frabit-server within daemon mode",
	Run: func(cmd *cobra.Command, args []string) {
		start()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start() {
	ctx := context.Background()
	cfg := config.Conf
	srv := server.NewServer(cfg)
	srv.Run(ctx)
}
