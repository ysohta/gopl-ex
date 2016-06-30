package tempconv

import (
	"fmt"
)

// Celsius represents temparature in the degree Celcius (°C).
type Celsius float64

// Fahrenheit represents temparature in the degree Fahrenheit (°F).
type Fahrenheit float64

// Kelvin represents temparature in the degree Kelvin (K).
type Kelvin float64

const (
	// AbsoluteZeroC represents value for the absolute zero temparature.
	AbsoluteZeroC Celsius = -273.15

	// FreezingC represents value for the water freezing point.
	FreezingC Celsius = 0

	// BoilingC represents value for the water boiling point.
	BoilingC Celsius = 100
)

// String returns Celsius degree with unit (°C).
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// String returns Fahrenheit degree with unit (°F).
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// String returns Kelvin degree with unit (°K).
func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
