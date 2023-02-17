package conv

import "math"

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  Celsius := (value - 32) * (5.0 / 9.0) // lagret svaret i egen variabel, som så blir omformet før den returneres
	return (math.Round(Celsius*100)/100) // ganger og deler med 100 for å får to desimaler, før svaret returneres
}

func CelsiusToFarhenheit(value float64) float64 {
  Farhenheit := (value * (9.0 / 5.0)) + 32
  return (math.Round(Farhenheit*100)/100)
}

func KelvinToCelsius(value float64) float64 {
  Celsius := (value - 273.15)
  return (math.Round(Celsius*100)/100)
}

func CelsiusToKelvin(value float64) float64 {
  Kelvin := (value + 273.15)
  return (math.Round(Kelvin*100)/100)
}

func KelvinToFarhenheit(value float64) float64 {
  Farhenheit := (value - 273.15) * (9.0 / 5.0) + 32
  return (math.Round(Farhenheit*100)/100)
}

func FarhenheitToKelvin(value float64) float64 {
  Kelvin := (value - 32) * (5.0 / 9.0) + 273.15
  return (math.Round(Kelvin*100)/100)
}