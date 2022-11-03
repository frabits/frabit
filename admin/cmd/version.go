/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"fmt"
	"github.com/frabit-io/frabit/common/version"
	"github.com/spf13/cobra"

	_ "github.com/frabit-io/frabit/common"
)

var newVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bytebase is a database schema change and version control tool",
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	v := version.Version{}
	fmt.Printf("%s\n", v.String())
}

func init() {
	rootCmd.AddCommand(newVersionCmd)
}
