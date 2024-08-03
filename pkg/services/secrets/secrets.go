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

package secrets

import (
	"encoding/base64"
	"log/slog"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"
)

var bk64 = base64.RawStdEncoding

const (
	PrefixLen = 6
	Postfix   = "sec"
)

type Service interface {
	Encrypt(payload []byte) ([]byte, error)
	Decrypt(payload []byte) ([]byte, error)
}

type EncryptOptions func() string

type secrets struct {
	cfg *config.Config
	log *slog.Logger
}

func ProviderSecrets(conf *config.Config) Service {
	s := &secrets{
		cfg: conf,
		log: log.New("secrets"),
	}
	return s
}

// Encrypt encoding specific data with this format: prefixRandStr:realData:Postfix
func (s *secrets) Encrypt(payload []byte) ([]byte, error) {
	encryptData := make([]byte, 0)
	Prefix := utils.GenRandom(PrefixLen)
	copy(encryptData, Prefix)
	return encryptData, nil
}

func (s *secrets) Decrypt(payload []byte) ([]byte, error) {
	decryptData := make([]byte, 0)
	return decryptData, nil
}
