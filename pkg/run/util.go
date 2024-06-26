/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package run

import "context"

// S is a short-hand for converting string slice syntaxes.
func S(s ...string) []string {
	return s
}

func getContextValue[T any](ctx context.Context, key any) T {
	x := ctx.Value(key)
	d, ok := x.(T)
	if !ok {
		var zero T
		return zero
	}
	return d
}
