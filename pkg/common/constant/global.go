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

package constant

type PrivilegeLevel string
type Role string
type AuditEvent string
type FeatureType string
type License string
type DbType string
type Theme string

const (
	GLOBAL    PrivilegeLevel = "global"
	WORKSPACE PrivilegeLevel = "workspace"
	PROJECT   PrivilegeLevel = "project"
	DATABASE  PrivilegeLevel = "database"
	GENERAl   PrivilegeLevel = "general"

	ADMIN  Role = "admin"
	EDITOR Role = "editor"
	VIEWER Role = "viewer"
	PTOSC       = "pt-online-schema-change"

	MainOrg     = "Main.org"
	GlobalOrgId = 1

	FeatureAuditLog  FeatureType = "fb.feature.audit-log"
	FeatureWatermark FeatureType = "fb.feature.water-mark"
	Community        License     = "community"
	Enterprise       License     = "enterprise"
	Ultimate         License     = "ultimate"

	Dark  Theme = "dark"
	Light Theme = "light"

	Mysql DbType = "mysql"
)
