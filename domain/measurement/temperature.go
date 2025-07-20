package measurement

import "fmt"

type TemperatureUnit int

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
	Kelvin
)

var _ Measurable = (*Temperature)(nil)

type Temperature struct {
	value float64
	unit  TemperatureUnit
}

func NewTemperature(value float64, unit TemperatureUnit) Temperature {
	return Temperature{
		value: value,
		unit:  unit,
	}
}

func (t Temperature) Celsius() float64 {
	switch t.unit {
	case Celsius:
		return t.value
	case Fahrenheit:
		return (t.value - 32) * 5 / 9
	case Kelvin:
		return t.value - 273.15
	default:
		return 0
	}
}

func (t Temperature) Fahrenheit() float64 {
	switch t.unit {
	case Celsius:
		return (t.value * 9 / 5) + 32
	case Fahrenheit:
		return t.value
	case Kelvin:
		return (t.value-273.15)*9/5 + 32
	default:
		return 0

	}
}
func (t Temperature) Kelvin() float64 {
	switch t.unit {
	case Celsius:
		return t.value + 273.15
	case Fahrenheit:
		return (t.value-32)*5/9 + 273.15
	case Kelvin:
		return t.value
	default:
		return 0
	}
}

func (t Temperature) String() string {
	switch t.unit {
	case Celsius:
		return fmt.Sprintf("%.2fC", t.Celsius())
	case Fahrenheit:
		return fmt.Sprintf("%.2fF", t.Fahrenheit())
	case Kelvin:
		return fmt.Sprintf("%.2fK", t.Kelvin())
	default:
		return fmt.Sprintf("%.2f", t.value)
	}
}

func (t Temperature) Equal(other Measurable) bool {
	u, ok := other.(Temperature)
	if !ok {
		return false
	}
	return t.value == u.value && t.unit == u.unit
}
