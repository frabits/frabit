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

const (
	GLOBAL    PrivilegeLevel = "global"
	WORKSPACE PrivilegeLevel = "workspace"
	PROJECT   PrivilegeLevel = "project"
	DATABASE  PrivilegeLevel = "database"
	GENERAl   PrivilegeLevel = "general"

	ADMIN  Role = "admin"
	EDITOR Role = "editor"
	VIEWER Role = "viewer"
	// FRABIT is the default metadata schema name if not provide a dbname via config file
	FRABIT = "frabit"
	PTOSC  = "pt-online-schema-change"

	CREATE_DATABASE AuditEvent = "Create database"

	FeatureAuditLog  FeatureType = "fb.feature.audit-log"
	FeatureWatermark FeatureType = "fb.feature.water-mark"
)
