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

package version

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/google/go-github/v50/github"
)

// initial version
const (
	Major = "1"
	Minor = "0"
	Patch = "0"
	Dist  = "community"
)

var (
	version   = "source"
	buildDate = "unknown"
	commit    = "unknown"
)

type Info struct {
	Version
	Build
}

type Version struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
	Dist  string `json:"dist"`
}

type Build struct {
	GitHash string `json:"git_hash"`
	Date    string `json:"date"`
	Arch    string `json:"arch"`
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
		Date:    buildDate,
		Arch:    runtime.GOARCH,
	}
}

func (v Version) versionString() string {
	var versionStr string
	if v.Dist == "" {
		versionStr = fmt.Sprintf("Version: %s.%s.%s", v.Major, v.Minor, v.Patch)
	} else {
		versionStr = fmt.Sprintf("Version: %s.%s.%s-%s", v.Major, v.Minor, v.Patch, v.Dist)
	}
	return versionStr
}

func (b Build) buildString() string {
	str := fmt.Sprintf("Hash: %s\nDate: %s\nArch: %s", b.GitHash, b.Date, b.Arch)
	return str
}

// String display frabit version and build information
func (info Info) String() string {
	str := fmt.Sprintf("%s\n%s\n", info.versionString(), info.buildString())
	return str
}

func getLatestVersion() string {
	ghClient := github.NewClient(nil)
	repo, _, err := ghClient.Repositories.GetLatestRelease(context.Background(), "frabits", "frabit")
	if err != nil {
		fmt.Errorf("cannot get latest version:%s", err)
		return ""
	}
	return *repo.TagName
}

func needToUpgrade(version, latest string) bool {
	return latest != "" && (strings.TrimPrefix(latest, "v") != strings.TrimPrefix(version, "v"))
}

func CheckLatestVersion() {
	fmt.Println("Check new version...")
	if version != "source" {
		latest := getLatestVersion()

		if ok := needToUpgrade(version, latest); ok {
			fmt.Printf("A newer version of the frabit is available,please upgrade to: %s\n", latest)
		}
	}

}
