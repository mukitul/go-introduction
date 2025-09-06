package main

import "fmt"

func main() {
	fmt.Print("Input temp in Fahrenheit: ")

	var tempInFahrenheit float64
	fmt.Scanf("%f", &tempInFahrenheit)

	tempInCelsius := convertFahrenheitToCelsius(tempInFahrenheit)

	fmt.Println("Temp in Celsius is: ", tempInCelsius)
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
