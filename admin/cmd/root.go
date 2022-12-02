/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	_ "github.com/frabits/frabit/common"
)

var rootCmd = &cobra.Command{
	Use:   "frabit-admin",
	Short: "Frabit is a comprehensive database platform for DBAs and developers",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(runtime.GOOS, runtime.GOARCH)
	},
}

var flag struct {
	port int
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// In the release build, Bytebase bundles frontend and backend together and runs on a single port as a mono server.
	// During development, Bytebase frontend runs on a separate port.
	rootCmd.PersistentFlags().IntVar(&flag.port, "port", 80, "port where Bytebase server runs. Default to 80")
}
