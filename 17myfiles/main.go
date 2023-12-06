package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"os"
)

func main() {
	
	fmt.Println("Wlecome to files of Golang")

	content := "This is my file and I am saving it to ./myfile.txt"

	file, err := os.Create("./myfiles.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content);

	if err != nil {
		panic(err)
	}

	fmt.Println("The length is:", length);
	defer file.Close()
	readFile("./myfiles.txt")
	
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("text data inside file is:", string(databyte));
}

func checkNilErr(err error) {
	if err!= nil {
		panic(err)
	}
}