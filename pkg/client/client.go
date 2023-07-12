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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	FrabitClient
}

// NewClient create a http client that used to make request to frabit-server
func NewClient(addr string) *Client {
	return &Client{
		FrabitClient: &frabitClient{addr: addr},
	}
}

type FrabitClient interface {
	GetVersion(ctx context.Context) (string, error)
}

type frabitClient struct {
	addr string
}

func (fc *frabitClient) GetVersion(ctx context.Context) (string, error) {
	endpoint := fmt.Sprintf("%s/admin/api/v1/version", fc.addr)
	resp, err := fc.doRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	var ver string
	err = unmarshalFrabitAPIRespone(resp.Body, &ver)
	return ver, err
}

// doRequest take a request to frabit-server
func (fc *frabitClient) doRequest(ctx context.Context, method string, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func unmarshalFrabitAPIRespone(r io.ReadCloser, v interface{}) error {
	defer func() {
		_ = r.Close()
	}()

	resp := struct {
		Status string          `json:"status"`
		Data   json.RawMessage `json:"data"`
	}{}

	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		fmt.Errorf("could not read respone: %w", err)
	}

	if v != nil && resp.Status == "success" {
		err := json.Unmarshal(resp.Data, v)
		if err != nil {
			return fmt.Errorf("unmarshaling respone: %w", err)
		}
	} else if resp.Status == "error" {
		return fmt.Errorf("unmarshaling respone: %w", err)
	}

	if resp.Status != "success" && resp.Status != "error" {
		return fmt.Errorf("unknown API respone status : %s", resp.Status)
	}
	return nil
}
