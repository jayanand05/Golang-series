package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops of golang");

	days := []string{"Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	
	// for i:= 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for i:= range days {
		fmt.Println(days[i])
	}
}