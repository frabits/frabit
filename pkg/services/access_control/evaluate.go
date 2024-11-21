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

package access_control

import (
	"context"
	"errors"
)

type Evaluator interface {
	Evaluate(ctx context.Context, permissions map[string][]string) (bool, error)
	MutateScopes(ctx context.Context, mutate ScopeAttributeMutator) (Evaluator, error)
}

type ScopeAttributeMutator func(context.Context, string) ([]string, error)

type evaluatorImpl struct {
	Action string
	Scopes []string
}

func EvalPermission(action string, scopes ...string) Evaluator {
	return &evaluatorImpl{Action: action, Scopes: scopes}
}

func (ei *evaluatorImpl) Evaluate(ctx context.Context, permissions map[string][]string) (bool, error) {
	return true, nil
}

func (ei *evaluatorImpl) MutateScopes(ctx context.Context, mutate ScopeAttributeMutator) (Evaluator, error) {
	scopes := make([]string, 0, len(ei.Scopes))
	if ei.Scopes == nil {
		return EvalPermission(ei.Action), nil
	}
	mutated := false
	for _, scope := range ei.Scopes {
		mutates, err := mutate(ctx, scope)
		if err != nil {
			return nil, err
		}
		mutated = true
		scopes = append(scopes, mutates...)
	}
	if !mutated {
		return nil, ErrScopeNotMutated

	}
	return EvalPermission(ei.Action, scopes...), nil
}

func EvaluateAny(evals ...Evaluator) Evaluator {
	return &evaluateAny{anyOf: evals}
}

type evaluateAny struct {
	anyOf []Evaluator
}

func (e *evaluateAny) Evaluate(ctx context.Context, permissions map[string][]string) (bool, error) {
	for _, eval := range e.anyOf {
		if ok, _ := eval.Evaluate(ctx, permissions); ok {
			return true, nil
		}
	}
	return false, nil
}

func (e *evaluateAny) MutateScopes(ctx context.Context, mutate ScopeAttributeMutator) (Evaluator, error) {
	evaluators := make([]Evaluator, len(e.anyOf))
	mutated := false
	for idx, eval := range e.anyOf {
		evaluator, err := eval.MutateScopes(ctx, mutate)
		if err != nil {
			if errors.Is(err, ErrScopeNotMutated) {
				evaluators[idx] = eval
				continue
			}
			return nil, ErrScopeNotMutated
		}
		mutated = true
		evaluators[idx] = evaluator
	}
	if !mutated {

		return nil, ErrScopeNotMutated
	}
	return EvaluateAny(evaluators...), nil
}

func EvaluateAll(evals ...Evaluator) Evaluator {
	return &evaluateAll{allOf: evals}
}

type evaluateAll struct {
	allOf []Evaluator
}

func (e *evaluateAll) Evaluate(ctx context.Context, permissions map[string][]string) (bool, error) {
	for _, eval := range e.allOf {
		if ok, err := eval.Evaluate(ctx, permissions); !ok {
			return false, err
		}
	}
	return true, nil
}

func (e *evaluateAll) MutateScopes(ctx context.Context, mutate ScopeAttributeMutator) (Evaluator, error) {
	evaluators := make([]Evaluator, len(e.allOf))
	mutated := false
	for idx, eval := range e.allOf {
		evaluator, err := eval.MutateScopes(ctx, mutate)
		if err != nil {
			if errors.Is(err, ErrScopeNotMutated) {
				evaluators[idx] = eval
				continue
			}
			return nil, ErrScopeNotMutated
		}

		mutated = true
		evaluators[idx] = evaluator
	}
	if !mutated {

		return nil, ErrScopeNotMutated
	}
	return EvaluateAll(evaluators...), nil
}
