package measurement

import (
	"github.com/psurti/stdxtra/domain/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {
	t.Run("NewTemperature", func(t *testing.T) {
		temp := NewTemperature(25.0, Celsius)
		assert.Equal(t, 25.0, temp.value)
		assert.Equal(t, Celsius, temp.unit)
	})

	t.Run("Celsius", func(t *testing.T) {
		t.Run("FromCelsius", func(t *testing.T) {
			temp := NewTemperature(25.0, Celsius)
			assert.Equal(t, 25.0, temp.Celsius())
		})

		t.Run("FromFahrenheit", func(t *testing.T) {
			temp := NewTemperature(77.0, Fahrenheit)
			assert.InDelta(t, 25.0, temp.Celsius(), 0.01)
		})

		t.Run("FromKelvin", func(t *testing.T) {
			temp := NewTemperature(298.15, Kelvin)
			assert.InDelta(t, 25.0, temp.Celsius(), 0.01)
		})
	})

	t.Run("Fahrenheit", func(t *testing.T) {
		t.Run("FromCelsius", func(t *testing.T) {
			temp := NewTemperature(25.0, Celsius)
			assert.InDelta(t, 77.0, temp.Fahrenheit(), 0.01)
		})

		t.Run("FromFahrenheit", func(t *testing.T) {
			temp := NewTemperature(77.0, Fahrenheit)
			assert.Equal(t, 77.0, temp.Fahrenheit())
		})

		t.Run("FromKelvin", func(t *testing.T) {
			temp := NewTemperature(298.15, Kelvin)
			assert.InDelta(t, 77.0, temp.Fahrenheit(), 0.01)
		})
	})

	t.Run("Kelvin", func(t *testing.T) {
		t.Run("FromCelsius", func(t *testing.T) {
			temp := NewTemperature(25.0, Celsius)
			assert.InDelta(t, 298.15, temp.Kelvin(), 0.01)
		})

		t.Run("FromFahrenheit", func(t *testing.T) {
			temp := NewTemperature(77.0, Fahrenheit)
			assert.InDelta(t, 298.15, temp.Kelvin(), 0.01)
		})

		t.Run("FromKelvin", func(t *testing.T) {
			temp := NewTemperature(298.15, Kelvin)
			assert.Equal(t, 298.15, temp.Kelvin())
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("Celsius", func(t *testing.T) {
			temp := NewTemperature(25.0, Celsius)
			assert.Equal(t, "25.00C", temp.String())
		})

		t.Run("Fahrenheit", func(t *testing.T) {
			temp := NewTemperature(77.0, Fahrenheit)
			assert.Equal(t, "77.00F", temp.String())
		})

		t.Run("Kelvin", func(t *testing.T) {
			temp := NewTemperature(298.15, Kelvin)
			assert.Equal(t, "298.15K", temp.String())
		})
	})

	t.Run("Equal", func(t *testing.T) {
		t.Run("SameValues", func(t *testing.T) {
			temp1 := NewTemperature(25.0, Celsius)
			temp2 := NewTemperature(25.0, Celsius)
			assert.True(t, temp1.Equal(temp2))
		})

		t.Run("DifferentValues", func(t *testing.T) {
			temp1 := NewTemperature(25.0, Celsius)
			temp2 := NewTemperature(30.0, Celsius)
			assert.False(t, temp1.Equal(temp2))
		})

		t.Run("DifferentUnits", func(t *testing.T) {
			temp1 := NewTemperature(25.0, Celsius)
			temp2 := NewTemperature(25.0, Fahrenheit)
			assert.False(t, temp1.Equal(temp2))
		})

		t.Run("DifferentValuesAndUnits", func(t *testing.T) {
			temp1 := NewTemperature(25.0, Celsius)
			temp2 := NewTemperature(77.0, Fahrenheit)
			assert.False(t, temp1.Equal(temp2))
		})
	})

	t.Run("EdgeCases", func(t *testing.T) {
		t.Run("AbsoluteZero", func(t *testing.T) {
			temp := NewTemperature(0, Kelvin)
			assert.InDelta(t, -273.15, temp.Celsius(), 0.01)
			assert.InDelta(t, -459.67, temp.Fahrenheit(), 0.01)
		})

		t.Run("FreezingPoint", func(t *testing.T) {
			temp := NewTemperature(0, Celsius)
			assert.InDelta(t, 32.0, temp.Fahrenheit(), 0.01)
			assert.InDelta(t, 273.15, temp.Kelvin(), 0.01)
		})

		t.Run("BoilingPoint", func(t *testing.T) {
			temp := NewTemperature(100, Celsius)
			assert.InDelta(t, 212.0, temp.Fahrenheit(), 0.01)
			assert.InDelta(t, 373.15, temp.Kelvin(), 0.01)
		})
	})

	t.Run("Values", func(t *testing.T) {
		var vals value.Values[Measurable]
		vals = append(vals, *value.New[Measurable](NewTemperature(2.0, Celsius)))
		vals = append(vals, *value.New[Measurable](NewTemperature(3.0, Celsius)))
	})
}
