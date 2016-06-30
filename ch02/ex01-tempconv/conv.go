package tempconv

// CToF converts Celsius to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToF converts Kelvin to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

// KToC converts Kelvin to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK converts Celsius to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToK converts Fahrenheit to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
