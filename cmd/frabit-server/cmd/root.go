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

	"github.com/frabits/frabit/pkg/common/version"
	"github.com/frabits/frabit/pkg/server"

	"github.com/spf13/cobra"
)

var StartOpts struct {
	Config  string
	Version bool
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "frabit-server",
	Short:   "Frabit The next-generation database automatic platform",
	Version: version.InfoStr.String(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if StartOpts.Version {
		fmt.Println(version.InfoStr.String())
		os.Exit(0)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start frabit-server",
	RunE:  start,
}

func start(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	srv, err := server.Initialize()
	if err != nil {
		return err
	}

	go listenToSystemSignals(ctx, srv)
	return srv.Run()
}

func listenToSystemSignals(ctx context.Context, src *server.Server) {
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
			if err := src.Shutdown(ctx, fmt.Sprintf("System signal:%s", sig)); err != nil {
				fmt.Fprintf(os.Stderr, "Timed out waiting for server to shut down")
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.Flags().BoolVar(&StartOpts.Version, "v", false, "Display frabit-server version")
	rootCmd.Flags().String(StartOpts.Config, "/etc/frabit/frabit-server.yml", "Specific frabit config file")
}
