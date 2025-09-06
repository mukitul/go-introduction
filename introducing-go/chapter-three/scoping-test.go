package main

import "fmt"

var x = "this is a package level variable"

func main() {
	fmt.Println("from main: ", x)
	anotherFunction()
}

func anotherFunction() {
	fmt.Println("from another function", x)
}
