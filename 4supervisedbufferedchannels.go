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
	deadLetterMessages := make(chan string, numberGoroutines)

	wg.Add(numberGoroutines + 1) //+1 because of supervisor

	go supervisor(deadLetterMessages, numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr, deadLetterMessages)
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

func worker(tasks chan string, worker int, deadLetter chan string) {
	defer func () {
		deadLetter <- fmt.Sprintf("Worked %d gone.", worker)
		wg.Done()
	}()

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

func supervisor(deadLetterMessages chan string, totalNumberOfWorkers int) {
	defer wg.Done()
	fmt.Printf("Supervisor started, waiting for: %d workers to die\n", totalNumberOfWorkers)
	var workersAlive = totalNumberOfWorkers
	for {
		deadLetter, ok := <- deadLetterMessages

		if !ok {
			fmt.Println("unexpected error waiting for dead letters")
			return
		}

		fmt.Printf("\nSupervisor received dead letter from worker: \"%s\"\n", deadLetter)
		workersAlive --
		if(workersAlive == 0) {
			fmt.Println("\nSupervisor received all dead letters, exiting supervisor routine")
			return
		} else {
			fmt.Printf("\nWaiting for %d dead letters ", workersAlive)
		}
	}

}
