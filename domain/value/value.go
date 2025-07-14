// Copyright 2023 The stdxtra Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package value

type Value[T any] struct {
	val T
}

//go:noinline
func Get[T any](v *Value[T]) T {
	return v.val
}

//go:noinline
func Set[T any](v *Value[T], val T) {
	v.val = val
}

//go:noinline
func New[T any](val T) *Value[T] {
	return &Value[T]{
		val: val,
	}
}

//go:noinline
func (v *Value[T]) Get() T {
	return v.val
}

//go:noinline
func (v *Value[T]) Set(val T) {
	v.val = val
}
