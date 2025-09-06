package main

import "fmt"

func main() {
	fmt.Println("sum of 1 + 1 is ", 1+1)

	fmt.Println("sum of 98.01 + 1.99 is ", 98.1+1.99)

	fmt.Println("lenght of 'Hello World' is ", len("Hello World"))

	fmt.Println("First character of 'Hello World' is ", "Hello World"[0]) // ASCII value of 'H' is 72

	fmt.Println("Is 1 equal to 1.0? ", 1 == 1.0) // true because Go converts 1 to 1.0 before comparison, why? because both are numeric types

	fmt.Println("Is 'Hello' equal to 'hello'? ", "Hello" == "hello") // false because string comparison is case-sensitive

	fmt.Println("Is true equal to false? ", true == false) // false because both are different boolean values

	fmt.Println("Value of Pi is ", 3.1415926535897939) // print 3.141592653589794, becase of float precision limit, always prints 15 decimal places

	fmt.Println("A big number: ", 200000000000000000)

}
