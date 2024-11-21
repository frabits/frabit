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
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"log/slog"

	"github.com/frabits/frabit/pkg/common/utils"
	"github.com/frabits/frabit/pkg/config"
	"github.com/frabits/frabit/pkg/infra/log"

	"golang.org/x/crypto/pbkdf2"
)

var bk64 = base64.RawStdEncoding

const (
	SaltLen = 6
	AesCfb  = "aes-cfb"
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

func (s *secrets) Encrypt(payload []byte) ([]byte, error) {
	encryptData := make([]byte, SaltLen+aes.BlockSize+len(payload))
	salt := utils.GenRandom(SaltLen)
	encryptKey, err := KeyToBytes(s.cfg.Server.SecureKey, salt)
	if err != nil {
		s.log.Error("create encryptKey failed", "Error", err.Error())
		return encryptData, err
	}

	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.

	copy(encryptData[:SaltLen], salt)
	iv := encryptData[SaltLen : SaltLen+aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encryptData[SaltLen+aes.BlockSize:], payload)

	return encryptData, nil
}

func (s *secrets) Decrypt(payload []byte) ([]byte, error) {
	if len(payload) < SaltLen {
		return nil, errors.New("unable to compute salt")
	}

	salt := payload[:SaltLen]
	encryptKey, err := KeyToBytes(s.cfg.Server.SecureKey, string(salt))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return nil, err
	}
	return decrypt(block, payload)
}

func decrypt(block cipher.Block, payload []byte) ([]byte, error) {
	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.
	if len(payload) < aes.BlockSize {
		return nil, errors.New("payload too short")
	}

	iv := payload[SaltLen : SaltLen+aes.BlockSize]
	payload = payload[SaltLen+aes.BlockSize:]
	payloadDst := make([]byte, len(payload))

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(payloadDst, payload)
	return payloadDst, nil
}

// KeyToBytes key length needs to be 32 bytes
func KeyToBytes(secret, salt string) ([]byte, error) {
	return pbkdf2.Key([]byte(secret), []byte(salt), 10000, 32, sha256.New), nil
}
