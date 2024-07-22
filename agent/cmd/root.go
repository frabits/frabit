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
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/frabits/frabit/pkg/agent"
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
	path string
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
	rootCmd.PersistentFlags().IntVar(&flag.port, "port", 19180, "port. Default to 19180")
}

func runAgent(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	a := agent.New(nil)
	if err := a.RunAgent(ctx); err != nil {
		a.Log.Error("Agent failed", "Error", err.Error())
		os.Exit(1)
	}
	listenToSystemSignals(ctx, a)
	a.RunAgent(ctx)
}

func listenToSystemSignals(ctx context.Context, agent *agent.Agent) {
	signalChan := make(chan os.Signal, 1)
	sighupChan := make(chan os.Signal, 1)

	signal.Notify(sighupChan, syscall.SIGHUP)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sighupChan:
			fmt.Println("nothing,maybe reload config in the future")
		case sig := <-signalChan:
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			if err := agent.Shutdown(ctx, fmt.Sprintf("System signal:%s", sig)); err != nil {
				fmt.Fprintf(os.Stderr, "Timed out waiting for server to shut down")
			}
		}
	}
}
