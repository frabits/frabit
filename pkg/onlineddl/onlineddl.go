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

package onlineddl

type DDLType string

const (
	// Direct represent directly execute statemet via "alter table"
	Direct DDLType = "direct"
	// Native represent MySQL native online DDL
	Native DDLType = "native"
	// Ghost represent gh-ost to take none-lock schema change
	Ghost DDLType = "gh-ost"
	// PtOSC represent pt-online-schema-change which provided by Percona
	PtOSC DDLType = "pt-osc"
)

type OnlineDDL struct {
	Type DDLType
}
