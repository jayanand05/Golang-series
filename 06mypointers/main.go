package main

import "fmt"fefrefr

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of the ptr is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the ptr is ", ptr)
	fmt.Println("Value of the ptr is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of this ptr is", *ptr)
}
