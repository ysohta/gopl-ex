package main

import (
	"fmt"
)

// Pound represents weight in Pound (lb).
type Pound float64

// Kilogram represents weight in Kilogram (kg).
type Kilogram float64

// LbToKg converts Pound to Kilogram.
func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / 0.45359237)
}

// KgToLb converts Kilogram to Pound.
func KgToLb(kg Kilogram) Pound {
	return Pound(kg * 0.45359237)
}

// String returns Pound weight with unit (lb).
func (lb Pound) String() string {
	return fmt.Sprintf("%glb", lb)
}

// String returns Kilogram weight with unit (m).
func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}
