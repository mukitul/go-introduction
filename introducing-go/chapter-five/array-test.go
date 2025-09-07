package main

import "fmt"

func main() {
	var a [6]float64 = [6]float64{98.0, 93.3, 77.9, 82.234, 83.98}

	a[len(a)-1] = 100.89

	var total float64 = 0.0

	for i:=0; i<len(a); i++ {
		total += a[i]
	}

	fmt.Println("Total (loop variant one): ", total)


	total = 0.0

	for index, value := range a {
		fmt.Printf("Index: %d, Value: %.2f\n", index, value)
		total += value
	}

	fmt.Println("Total (loop variant two): ", total)


}
