package main

import "fmt"

func main() {
	var x string = "Hello world!"
	fmt.Println(x)

	y := "Hello Again!";
	fmt.Println(y)

	var name = "MUKITUL"
	fmt.Println(name)

	var age, height = 24, 5.7
	fmt.Println(age, height)

	var (
		country  = "Bangladesh"
		city     = "Dhaka"
		postCode = 1207
	)
	fmt.Println(country, city, postCode)
}