package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza store")
	fmt.Println("Enter your rating between 1 to 10");

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	}else {
		
	fmt.Println("Thanks for rating,", numrating + 4)
	}
}