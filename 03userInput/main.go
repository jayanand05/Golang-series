package main
gghjnjnj
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to our pizza store")

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter your value: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
}
