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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		v := New(42)
		assert.Equal(t, 42, v.Get())
	})

	t.Run("Set", func(t *testing.T) {
		v := New(0)
		v.Set(100)
		assert.Equal(t, 100, v.Get())
	})

	t.Run("Get", func(t *testing.T) {
		v := New("test")
		assert.Equal(t, "test", v.Get())
	})

	t.Run("Set (package)", func(t *testing.T) {
		var v Value[float64]
		Set(&v, 3.14)
		assert.Equal(t, 3.14, Get(&v))
	})
}
