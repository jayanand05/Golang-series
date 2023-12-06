package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct part of the golang");
  
	userinput := User{"Jay Anand", "jay01541@gmail.com", true, 22}
	fmt.Println(userinput)
	fmt.Printf("userinput details are : %+v\n", userinput)
	fmt.Printf("Name is %v and Email is %v. ", userinput.Name, userinput.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
} 