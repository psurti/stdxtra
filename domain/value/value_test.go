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
