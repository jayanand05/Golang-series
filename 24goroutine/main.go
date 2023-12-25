package main 

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	webList := []string {
		"https://google.com",
		"https://amazon.com",
		"https://apple.com",
	}
	for _, web := range webList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)

	if err!=nil {
		fmt.Println("OOPS system is not working")
	}
	fmt.Printf("%d is status code with %s\n", result.StatusCode, endpoint)
}
