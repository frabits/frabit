// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2024 Frabit Team
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

package satoken

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/frabits/frabit/pkg/config"
	"github.com/pkg/errors"
	"hash/crc32"
	"strings"

	"github.com/frabits/frabit/pkg/common/utils"
)

const (
	FrabitPrefix = "fb"
)

type KeyGenResult struct {
	HashedKey    string
	ClientSecret string
}

type PrefixToken struct {
	Secret   string
	Checksum string
}

var ErrInvalidApiKey = errors.New("invalid API key")

func New() *KeyGenResult {
	token := &KeyGenResult{}
	secret := utils.GenRandom(32)
	pt := PrefixToken{Secret: secret}
	pt.Checksum = pt.GenerateChecksum()
	token.HashedKey = pt.Hash()
	token.ClientSecret = pt.String()
	return token
}

func (pt *PrefixToken) key() string {
	return fmt.Sprintf("%s_%s", FrabitPrefix, pt.Secret)
}

func (pt *PrefixToken) String() string {
	return fmt.Sprintf("%s_%s", pt.key(), pt.Checksum)
}

func (pt *PrefixToken) Hash() string {
	hash := sha256.New()
	hash.Write([]byte(pt.Secret + config.Security))
	return hex.EncodeToString(hash.Sum(nil))
}

func (pt *PrefixToken) GenerateChecksum() string {
	checksum := crc32.ChecksumIEEE([]byte(pt.key()))
	//checksum to []byte
	checksumBytes := make([]byte, 4)
	checksumBytes[0] = byte(checksum)
	checksumBytes[1] = byte(checksum >> 8)
	checksumBytes[2] = byte(checksum >> 16)
	checksumBytes[3] = byte(checksum >> 24)

	return hex.EncodeToString(checksumBytes)
}

func Decode(keyString string) (*PrefixToken, error) {
	if !strings.HasPrefix(keyString, FrabitPrefix) {
		return nil, ErrInvalidApiKey
	}

	parts := strings.Split(keyString, "_")
	if len(parts) != 3 {
		return nil, ErrInvalidApiKey
	}

	key := &PrefixToken{
		Secret:   parts[1],
		Checksum: parts[2],
	}
	if key.GenerateChecksum() != key.Checksum {
		return nil, ErrInvalidApiKey
	}

	return key, nil
}
