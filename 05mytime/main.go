package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05  Monday  "))
}