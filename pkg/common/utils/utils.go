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

package utils

import (
	"archive/tar"
	"compress/gzip"
	"github.com/google/uuid"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/xi2/xz"
)

// ExtractTarGz extracts the given file as .tar.gz format to the given directory.
func ExtractTarGz(tarGzF io.Reader, targetDir string) error {
	gzipR, err := gzip.NewReader(tarGzF)
	if err != nil {
		return err
	}
	defer gzipR.Close()
	return extractTar(gzipR, targetDir)
}

// ExtractTarXz extracts the given file as .tar.xz or .txz format to the given directory.
func ExtractTarXz(tarXzF io.Reader, targetDir string) error {
	xzR, err := xz.NewReader(tarXzF, 0)
	if err != nil {
		return err
	}
	return extractTar(xzR, targetDir)
}

func extractTar(r io.Reader, targetDir string) error {
	tarReader := tar.NewReader(r)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath, err := filepath.Abs(filepath.Join(targetDir, header.Name))
		if err != nil {
			return errors.Wrapf(err, "failed to get absolute path for %q", header.Name)
		}

		// Ensure that output paths constructed from zip archive entries
		// are validated to prevent writing files to unexpected locations.
		if strings.Contains(targetPath, "..") {
			return errors.Errorf("invalid path %q", targetPath)
		}

		if err := os.MkdirAll(path.Dir(targetPath), os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to create directory %q", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeReg:
			outFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			var totalWritten int64
			for totalWritten < header.Size {
				written, err := io.CopyN(outFile, tarReader, 1024)
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}
				totalWritten += written
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, targetPath); err != nil {
				return err
			}
		default:
			return errors.Errorf("unsupported type flag %d", header.Typeflag)
		}
	}

	return nil
}

func createUUID() string {
	return uuid.NewString()
}

func CreateUUIDWithDelimiter(delimiter string) string {
	rawUUID := createUUID()
	return strings.Replace(rawUUID, "-", delimiter, 0)
}

func NowDatetime() string {
	return time.Now().Format(time.DateTime)
}
