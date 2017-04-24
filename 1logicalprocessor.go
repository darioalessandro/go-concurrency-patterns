package main

import (
	"runtime"
	"sync"
	"fmt"
	"time"
)

/*
Program that creates two goroutines that display the English alphabet with lower and uppercase letters in a concurrent
fashion.
Excerpt From: William Kennedy. “Go in Action.” iBooks.
 */

func main() {
	runtime.GOMAXPROCS(1) //limiting runtime to use 1 logical processor.
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func () {
		defer wg.Done() //defer get's called before the function returns
		fmt.Println("\nentering upper case")
		//Display the alphabet three times
		fmt.Println("\nupper case after timer")
		for count:= 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func () {
		defer wg.Done() //defer get's called before the function returns
		fmt.Println("\nentering lower case")
		//Display the alphabet three times
		fmt.Println("\nlower case after timer")
		for count:= 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

}