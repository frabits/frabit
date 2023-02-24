/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"context"

	"github.com/frabits/frabit/server"
	"github.com/frabits/frabit/server/config"
	"github.com/frabits/frabit/server/db"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start frabit-server within daemon mode",
	Run: func(cmd *cobra.Command, args []string) {
		start()

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func start() {
	ctx := context.Background()
	cfg := config.Config{}
	db, _ := db.OpenFrabit()
	db.Ping()
	srv := server.NewServer(cfg)
	srv.Run(ctx, 9180)
}
