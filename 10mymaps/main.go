package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps of golang")rfrfrf

	languages := make(map[string]int)
	languages["JS"] = 8
	languages["LI"] = 9
	languages["MI"] = 7
	languages["SI"] = 6
	languages["PI"] = 2
	fmt.Println(languages)
	fmt.Println("JS number is", languages["JS"])

	delete(languages, "LI")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("The key is %v , and the value is %v\n", key, value)
	}
}
