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
	"testing"

	"github.com/frabits/frabit/pkg/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Encrypt(t *testing.T) {
	svc := newSecretsForTest()
	toEncrypt := []byte("frabit platform is cool")
	encrypted, err := svc.Encrypt(toEncrypt)
	require.NoError(t, err)
	assert.NotNil(t, encrypted)
	assert.NotEmpty(t, encrypted)
}

func Test_Decrypt(t *testing.T) {
	svc := newSecretsForTest()
	toEncrypt := []byte("frabitisaprettycoolp135%latform")
	encrypted, err := svc.Encrypt(toEncrypt)

	encryptedStr := bk64.EncodeToString(encrypted)
	toDecrypt, _ := bk64.DecodeString(encryptedStr)
	decrypted, err := svc.Decrypt(toDecrypt)
	require.NoError(t, err)
	assert.NotNil(t, decrypted)
	assert.NotEmpty(t, decrypted)
}

func newSecretsForTest() *secrets {
	cfg := &config.Config{}
	cfg.Server.SecureKey = "a_SecureKey"
	return &secrets{cfg: cfg}
}
