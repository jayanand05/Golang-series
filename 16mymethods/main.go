package main

import "fmt"

func main() {
	fmt.Println("Wlecome to methods of golang")
	userInput := User{"Jay", "hi@gmail.com", true, 88}
	fmt.Println(userInput)
	fmt.Printf("All details are: %+v\n", userInput)
	fmt.Printf("Name is %v and status is %v\n", userInput.Name, userInput.Status)
	userInput.SetMy()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// func (u User) setCalp() {
// 	fmt.Println("Wleocme to m y word ", u.Email)
// }

func (u User) SetMy() {
	u. Email = "jay01541@gmail.com"
	fmt.Println("new gamil is ", u.Email)
}
