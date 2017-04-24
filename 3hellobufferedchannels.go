package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

const (
	numberGoroutines 	= 4
	taskLoad 		= 10
)

var wg sync.WaitGroup

//This function gets called before main
func init() {
	rand.Seed(time.Now().Unix())
}

func  main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//Add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//Close the channel so the goroutines will quit
	//when all the work is done.
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		//Wait for work to be posted
		task, ok := <- tasks
		if !ok {
			//This means the channel is
			//empty && closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		//Display we are starting the work
		fmt.Printf("Worker: %d: Started %s\n", worker, task)

		//Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}
