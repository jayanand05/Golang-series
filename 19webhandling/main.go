package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to web handling of golang")
    
     fmt.Println(myUrl)

	 result, _ := url.Parse(myUrl)
	 fmt.Println(result.Scheme)
	 fmt.Println(result.Port())
}
