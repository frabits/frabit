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

package config

import (
	"os"
)

type Config struct {
	service    string
	fbToken    string
	KeyringDir string
}

/*
service := "my-app"
user := "anon"
password := "secret"

// set password
err := keyring.Set(service, user, password)
if err != nil {
log.Fatal(err)
}

// get password
secret, err := keyring.Get(service, user)
if err != nil {
log.Fatal(err)
}
log.Println(secret)
*/

func (cfg *Config) HasTokenFromEnv() bool {

	return true
}

func (cfg *Config) HasTokenFromKeyring() bool {

	return true
}

func (cfg *Config) SetTokenToKeyring(token string) error {
	if err := keyring.Set(cfg.service, cfg.fbToken, token); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetTokenToEnv(token string) error {
	if err := os.Setenv(cfg.fbToken, token); err != nil {
		return err
	}
	return nil
}
