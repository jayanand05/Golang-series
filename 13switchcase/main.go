package main dedede

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switchcase part of the golang")
	
  diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)
}
