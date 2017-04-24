package main

import (
	"sync"
	"fmt"
)


/*

functionA   -> hello -> functionB

after functionB receives hello, the program will terminate.

 */

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	channel := make(chan string)

	var functionA = func (c chan string) {
		defer wg.Done()
		c <- "hello"
	}

	var functionB = func (c chan string) {
		defer wg.Done()
		message := <- c
		fmt.Printf("B: got message from A %s", message)
	}

	go functionA(channel)
	go functionB(channel)

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}