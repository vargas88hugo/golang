// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

// Celsius represents scale C
type Celsius float64

// Fahrenheit represents scale F
type Fahrenheit float64

// Constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

// CToF Function that converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Function that converts Fahrenhei to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
