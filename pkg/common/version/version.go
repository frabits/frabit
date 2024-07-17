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
	"github.com/briandowns/spinner"
	"runtime"
	"strings"
	"time"

	"github.com/google/go-github/v50/github"
)

// initial Version
const (
	Major = "1"
	Minor = "0"
	Patch = "0"
	Dist  = "community"
)

var (
	Version   = "source"
	BuildDate = "unknown"
	Commit    = "unknown"
)

type Info struct {
	Major   string `json:"major"`
	Minor   string `json:"minor"`
	Patch   string `json:"patch"`
	Dist    string `json:"dist"`
	GitHash string `json:"git_hash"`
	Date    string `json:"date"`
	Arch    string `json:"arch"`
}

var InfoStr Info

func init() {
	InfoStr = newInfo()
}

func newInfo() Info {
	var (
		major string
		minor string
		patch string
	)
	fmt.Printf("verison:%s BuildDate:%s Commit:%s\n", Version, BuildDate, Commit)
	if Version != "source" {
		Version = strings.TrimPrefix(Version, "v")

		major = strings.Split(Version, ".")[0]
		minor = strings.Split(Version, ".")[1]
		patch = strings.Split(Version, ".")[2]
	}
	return Info{
		Major:   major,
		Minor:   minor,
		Patch:   patch,
		Dist:    Dist,
		GitHash: Commit,
		Date:    BuildDate,
		Arch:    runtime.GOARCH,
	}
}

func (i Info) versionString() string {
	versionStr := fmt.Sprintf("Version: %s.%s.%s", i.Major, i.Minor, i.Patch)
	if i.Dist != "" {
		versionStr = fmt.Sprintf("%s-%s", versionStr, i.Dist)
	}
	return versionStr
}

func (i Info) buildString() string {
	return fmt.Sprintf("Commit: %s BuildDate:%s Arch:%s", i.GitHash, i.Date, i.Arch)
}

// String display frabit Version and build information
func (i Info) String() string {
	str := fmt.Sprintf("%s\n%s\n", i.versionString(), i.buildString())
	return str
}

func getLatestVersion() string {
	ghClient := github.NewClient(nil)
	repo, _, err := ghClient.Repositories.GetLatestRelease(context.Background(), "frabits", "frabit")
	if err != nil {
		fmt.Errorf("cannot get latest Version:%s", err)
		return ""
	}
	return *repo.TagName
}

func needToUpgrade(version, latest string) bool {
	return latest != "" && (strings.TrimPrefix(latest, "v") != strings.TrimPrefix(version, "v"))
}

func CheckLatestVersion() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	defer s.Stop()
	s.Start()
	s.Prefix = "Check new Version..."
	if Version != "source" {
		latest := getLatestVersion()

		if ok := needToUpgrade(Version, latest); ok {
			fmt.Printf("A newer Version of the frabit is available,please upgrade to: %s\n", latest)
		}
	}

}
