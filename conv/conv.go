package conv

var Celsius float64
var Kelvin float64
var Fahrenheit float64

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius = (Fahrenheit - 32) * (5 / 9)

	return (value - 32) * (5 / 9)
}

// Konverterer celsius til fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit = Celsius*(9/5) + 32

	return Fahrenheit
}

// Konverterer celsius til kelvin
func CelsiusToKelvin(value float64) float64 {
	Kelvin = Celsius + 273.15

	return Kelvin
}

// konverterer kelvin til celsius
func KelvinToCelsius(value float64) float64 {
	Celsius = Kelvin - 273.15

	return Celsius
}

// konverterer kelvin til fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit = (Kelvin-273.15)*(9/5) + 32

	return Fahrenheit
}

// konverterer fahrenheit til kelvin
func FahrenheitToKelvin(value float64) float64 {
	Kelvin = (Fahrenheit-32)*(5/9) + 273.15

	return Kelvin
}
