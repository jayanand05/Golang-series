package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello and welcome to web request in Golang")

	const url = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(url)

	if err!= nil {
		println(err)
	}

	fmt.Printf("Type of response is : %T\n", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err!= nil {
		println(err)
	}
	fmt.Println(string(databytes))
}
