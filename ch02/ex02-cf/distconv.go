package main

import (
	"fmt"
)

// Feet represents distance in Feet (ft).
type Feet float64

// Meter represents distance in Meter (m).
type Meter float64

// FtToM converts Feet to Meter.
func FtToM(ft Feet) Meter {
	return Meter(ft * 0.3048)
}

// MToFt converts Meter to Feet.
func MToFt(m Meter) Feet {
	return Feet(m / 0.3048)
}

// String returns Feet distance with unit (ft).
func (ft Feet) String() string {
	return fmt.Sprintf("%gft", ft)
}

// String returns Meter distance with unit (m).
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
