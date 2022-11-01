/*

  (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
  GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

  This file is part of Frabit

*/

package version

import (
	"fmt"
	"time"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Dist  string
}

type Build struct {
	GitHash string
	Date    time.Time
	Arch    string
}

func init() {
	newVersion()
}
func newVersion() *Version {
	return &Version{}

}

func (v Version) String() string {
	str := fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}

func (v Version) BuildInfo() string {
	str := fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}
