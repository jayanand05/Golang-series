package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else part of golang")
	loginCount := 44
	

	if loginCount < 10 {
		fmt.Println("login count is more than 10");

	}else if loginCount > 50 {
		fmt.Println("loging is more than 50")
	}else {
		fmt.Println("loging is unknown")
	}

	if 6%2 == 0 {
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}
}