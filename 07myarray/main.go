package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array of golang")

	var fruitList[4] string;

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Guava"
	fruitList[3] = "Litchi"
	fmt.Println("Fruitlist is :", fruitList)
	fmt.Println("Length of fruitlist is ", len(fruitList))

	var veggyList = [3]string{"Potato", "Onion", "Bringle"}
	fmt.Println("The list of veggy is ", veggyList)
	fmt.Println("The list of veggy is ", len(veggyList))

}
