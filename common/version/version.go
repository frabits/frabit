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

func (v Version) versionString() string {
	str := fmt.Sprintf("%s.%s.%s-%s", v.Major, v.Minor, v.Patch, v.Dist)
	return str
}

func (b Build) buildString() string {
	str := fmt.Sprintf("Hash: %s\nDate: %s\nArch: %s\n", b.GitHash, b.Date, b.Arch)
	return str
}

// String display frabit version and build information
func (info Info) String() string {
	str := fmt.Sprintf("%s\n%s\n", info.versionString(), info.buildString())
	return str
}
