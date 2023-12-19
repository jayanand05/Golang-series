package main

import "fmt"
cdcdcd
func main() {
	fmt.Println("HIIII")
	greeter()

	result := add(3, 5)
	fmt.Println("Result is ", result)

	proResult, my := proAdder(2, 3, 4, 5, 6, 710)
	fmt.Println("Result is", proResult)
	fmt.Println("message is ", my)
	myWall()
	
}

func myWall() {
	for i := 0; i < 5; i++{
		fmt.Print(i)
	}
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0;

	for _ , val := range values{
		total += val
	}
	return total, "hiiiieieyehhhhbhj"
}

func greeter() {
	fmt.Println("hiiiiiiiii")
	}
