package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice part of golang series ")frfrfr

	var fruitlist = []string{"Banana", "Guava", "Litchi"}
	fmt.Println(fruitlist)
	fmt.Printf("Type of fruilist is %T\n", fruitlist)

	var fruilist = append(fruitlist, "Mango", "Jackfruit")
	fmt.Println(fruilist)

	fruilist = append(fruilist[1:2])
	fmt.Println(fruilist)

	highScores := make([]int, 4)

	highScores[0] = 4456
	highScores[1] = 88
	highScores[2] = 44
	highScores[3] = 4
	fmt.Println(highScores)

	highScores = append(highScores, 987, 999)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"Java", "Python", "Ruby", "Rust", "Mongo", "React"}
	fmt.Println(courses)

	var index int = 5
	courses = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)
}
