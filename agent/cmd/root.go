/*
Copyright © 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"fmt"
	"os"

	"github.com/frabits/frabit/common/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "frabit-agent",
	Short:   "Frabit is a comprehensive database platform for DBAs and developers",
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
	fmt.Printf("%s", version.InfoStr.BuildInfo())
}
