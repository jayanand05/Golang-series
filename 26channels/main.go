package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channles in Golang")

	mych := make(chan int, 3)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func (ch chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <- mych
		fmt.Println(isChannelOpen)
		fmt.Println(value)
		// fmt.Println(<-mych)
		
		wg.Done()
	} (mych, wg) 
	go func (ch chan int, wg *sync.WaitGroup) {
		mych <- 0
		close(mych)
		
	
		wg.Done()
	} (mych, wg) 
	wg.Wait()
} 