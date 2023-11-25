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

package cmdutil

import (
	"github.com/frabits/frabit/pkg/common/config"

	"github.com/spf13/cobra"
)

// DisableAuthCheck disable a command authority
func DisableAuthCheck(cmd *cobra.Command) {
	if cmd.Annotations == nil {
		cmd.Annotations = map[string]string{}
	}

	cmd.Annotations["NeedAuthCheck"] = "false"
}

// CheckAuth if a command already been auth via a token
func CheckAuth(cfg config.Config) bool {
	if ok := cfg.HasTokenFromEnv(); ok {
		return true
	} else if ok := cfg.HasTokenFromKeyring(); ok {
		return true
	}
	return false
}

// IsAuthCheckEnabled ensure a command already
func IsAuthCheckEnabled(cmd *cobra.Command) bool {
	switch cmd.Name() {
	case "help", "version":
		return false
	}
	for c := cmd; c.Parent() != nil; c = c.Parent() {
		if c.Annotations != nil && c.Annotations["NeedAuthCheck"] == "false" {
			return false
		}
	}
	return true
}
