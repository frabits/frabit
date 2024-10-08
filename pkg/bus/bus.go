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

package bus

import "context"

type HandlerFunc any

type Bus interface {
	Publish(ctx context.Context, Msg any) error
	AddListener(ctx context.Context, listener HandlerFunc) error
}

type InProBus struct {
	Listener map[string][]HandlerFunc
	BusCh    chan any
}

func ProviderBus() Bus {
	return &InProBus{}
}

func (i *InProBus) Publish(ctx context.Context, Msg any) error {
	for {
		select {
		case <-ctx.Done():
		default:
			i.BusCh <- Msg
		}
	}
}

func (i *InProBus) AddListener(ctx context.Context, listener HandlerFunc) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:

			return nil
		}
	}
}
