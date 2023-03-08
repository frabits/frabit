/*
Copyright © 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package version

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// initial version
const (
	Major = "1"
	Minor = "0"
	Patch = "0"
	Dist  = "community"
)

var (
	version = "source"
	commit  = ""
)

type Info struct {
	Version
	Build
}

type Version struct {
	Major string
	Minor string
	Patch string
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
	if version != "source" {
		version = strings.TrimPrefix(version, "v")
		return Version{
			Major: strings.Split(version, ".")[0],
			Minor: strings.Split(version, ".")[1],
			Patch: strings.Split(version, ".")[2],
			Dist:  Dist,
		}
	}

	return Version{
		Major: Major,
		Minor: Minor,
		Patch: Patch,
		Dist:  Dist,
	}
}

func newBuild() Build {
	return Build{
		GitHash: commit,
		Date:    time.Now().Format(time.RFC3339),
		Arch:    runtime.GOARCH,
	}
}

func (v Version) String() string {
	str := fmt.Sprintf("%s.%s.%s-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}

func (b Build) BuildInfo() string {
	str := fmt.Sprintf("Hash: %s\nDate: %s\nArch: %s\n", b.GitHash, b.Date, b.Arch)
	return str
}
