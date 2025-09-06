package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {

		fmt.Println(i)
		i++
	}


	anotherVariation()

	switchVariation()
	
}

func switchVariation() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2 or 3")
	}	
}

func anotherVariation() {
	for i:=1; i<=10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}