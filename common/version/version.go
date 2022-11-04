/*

  (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
  GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

  This file is part of Frabit

*/

package version

import (
	"fmt"
	"runtime"
	"time"
)

const (
	Major = 1
	Minor = 0
	Patch = 0
	Dist  = "community"
)

type Info struct {
	Version
	Build
}

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

var InfoStr Info

func init() {
	InfoStr = newInfo()
}

func newInfo() Info {
	return Info{
		Version: newVersion(),
		Build:   newBuild(),
	}
}
func newVersion() Version {
	return Version{
		Major: Major,
		Minor: Minor,
		Patch: Patch,
		Dist:  Dist,
	}
}

func newBuild() Build {
	return Build{
		GitHash: "unknown",
		Date:    time.Now().Format(time.RFC3339),
		Arch:    runtime.GOARCH,
	}
}

func (v Version) String() string {
	str := fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}

func (b Build) BuildInfo() string {
	str := fmt.Sprintf("GitHash:%s BuildDate:%s Arch:%s", b.GitHash, b.Date, b.Arch)
	return str
}
