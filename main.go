package main

import (
	"flag"
	"fmt"

	"github.com/mathiasirgemo/funtemps/conv"
	"github.com/mathiasirgemo/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var Fahrenheit float64
var Celsius float64
var Kelvin float64

var out string

var funfact string

var Luna string
var Sun string
var Terra string
var t string


// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {


	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&Fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&Kelvin, "K", 1, "temperatur i grader Kelvin")
	flag.Float64Var(&Celsius, "C", 0.0, "temperatur i grader Celsius")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "temperaturskala når funfacts skal vises")


}

//gir output i kommandolinje
func main() {
	flag.Parse()


	if isFlagPassed("F") && out == "C" {
		Celsius = conv.FahrenheitToCelsius(Fahrenheit)
		fmt.Println(Fahrenheit, "°F er", Celsius, "°C")
	} else if isFlagPassed("C") && out == "F" {
		Fahrenheit = conv.CelsiusToFahrenheit(Celsius)
		fmt.Println(Celsius, "°C er", Fahrenheit, "°F")
	} else if isFlagPassed("C") && out == "K" {
		Kelvin = conv.CelsiusToKelvin(Celsius)
		fmt.Println(Celsius, "°C er", Kelvin, "K")
	} else if isFlagPassed("K") && out == "C" {
		Celsius = conv.KelvinToCelsius(Kelvin)
		fmt.Println(Kelvin, "K er", Celsius, "°C")
	} else if isFlagPassed("K") && out == "F" {
		Fahrenheit = conv.KelvinToFahrenheit(Kelvin)
		fmt.Println(Kelvin, "K er", Fahrenheit, "°F")
	} else if isFlagPassed("F") && out == "K" {
		Kelvin = conv.FahrenheitToKelvin(Fahrenheit)
		fmt.Println(Fahrenheit, "°F er", Kelvin, "K")
	} else if isFlagPassed("funfacts") && isFlagPassed("t") && t == "C" {
		facts := funfacts.GetFunFacts(funfact)
		for i, fact := range facts {
			fmt.Printf("%d: %s\n", i+1, fact)
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}