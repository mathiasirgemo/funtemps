package conv

var Celsius float64
var Kelvin float64
var Fahrenheit float64

// Konverterer Farhenheit til Celsius //OK
func FahrenheitToCelsius(value float64) float64 {

	return (value - 32.0) * 5.0 / 9.0
}

// Konverterer celsius til fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit := value*(9.0/5.0) + 32.0
	return Fahrenheit
}

// Konverterer celsius til kelvin
func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

// konverterer kelvin til celsius
func KelvinToCelsius(value float64) float64 {
	Celsius = Kelvin - 273.15

	return value - 273.15
}

// konverterer kelvin til fahrenheit //RIKTIG
func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit = (Kelvin * 9.0 / 5.0) - 460.0

	return (value * 9.0 / 5.0) - 460.0
}

// konverterer fahrenheit til kelvin //FEIL
func FahrenheitToKelvin(value float64) float64 {
	Kelvin = (Fahrenheit + 460.0) * 5.0 / 9.0

	return (value + 460.0) * 5.0 / 9.0
}
