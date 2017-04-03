package main

import "fmt"

var WorkerQueue chan chan WorkRequest //buffered channel of channel

func StartDispatcher(nworkers int) {
	//first, initialize the channel
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	//now , create all of our workers
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting Worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
