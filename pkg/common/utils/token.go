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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	mrand "math/rand"
	"time"
)

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", nil
	}
	return hex.EncodeToString(bytes), nil
}

func GenRandom(num int) string {
	metaStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Str := []byte(metaStr)
	var targetStr []byte
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		targetStr = append(targetStr, Str[r.Intn(len(Str))])
	}
	return string(targetStr)

}

type Token struct {
	Hash string
}

func NewToken(n int) *Token {
	newHash := GenRandom(n)
	return &Token{
		Hash: newHash,
	}
}

func GenHash(publicAuth, privateAuth string) string {
	hashRaw := fmt.Sprintf("%s:%s", publicAuth, privateAuth)
	hash, err := bcrypt.GenerateFromPassword([]byte(hashRaw), 32)
	if err != nil {
		return ""
	}
	return string(hash)
}
