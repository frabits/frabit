/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/
package db

// generateSQLPatches contains DDLs for patching schema to the latest version.
// Add new statements at the end of the list so they form a changelog.
var generateSQLPatches = []string{}
