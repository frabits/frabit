/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"fmt"
	"github.com/frabits/frabit/common/version"
	"github.com/spf13/cobra"

	_ "github.com/frabits/frabit/common"
)

var newVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Frabit is a comprehensive database platform for DBAs and developers",
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
	fmt.Printf("%s", version.InfoStr.BuildInfo())
}

func init() {
	rootCmd.AddCommand(newVersionCmd)
}
