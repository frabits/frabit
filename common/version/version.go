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

const (
	Major = 1
	Minor = 0
	Patch = 0
	Dist  = "community"
)

type Version struct {
	Major int
	Minor int
	Patch int
	Dist  string
}

type Build struct {
	GitHash string
	Date    string
	Arch    string
}

func init() {
	newVersion()
	newBuild()
}

func newVersion() *Version {
	return &Version{
		Major: Major,
		Minor: Minor,
		Patch: Patch,
		Dist:  Dist,
	}
}

func newBuild() *Build {
	return &Build{
		GitHash: "unknown",
		Date:    time.Now().Format(time.RFC3339),
		Arch:    "",
	}
}

func (v Version) String() string {
	str := fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}

func (b Build) BuildInfo() string {
	str := fmt.Sprintf("GitHash:%s\n Date:%s\n Arch:%s\n", b.GitHash, b.Date, b.Arch)
	return str
}
