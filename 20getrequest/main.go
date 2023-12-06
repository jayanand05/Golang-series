package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to getRequest of golang")
    //performGetRequest()
	performPostRequest()
}

func performGetRequest() {
  const myurl = "http://localhost:8000/get"

  response, err := http.Get(myurl)
  if err != nil {
	panic(err)
  }

  defer response.Body.Close()
  fmt.Printf("Type of response is %T\n:", response)
  fmt.Println("status code is :", response.StatusCode)
  fmt.Println("Content length is:", response.ContentLength);

  databytes, err := ioutil.ReadAll(response.Body)
  if err != nil {
	panic(err)
  }
  fmt.Println(string(databytes))
}

func performPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	  {
		"name" : "Jay",
		"course": "Blockchain",
		"age" : 22
	  }
	`)

	response , err := http.Post(myUrl, "application/json", requestBody)
	if err!= nil {
		panic(err)
	}

	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err!= nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}