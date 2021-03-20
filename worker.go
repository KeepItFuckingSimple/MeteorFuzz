package main

import (
	"fmt"
	"net/http"
)

type Worker struct {
	requests []*http.Request
}

type WorkerPool struct {
	workers []Worker
	results []string
}

func (w *Worker) Run() []string{ //SAMPLE WORKER CODE SIMULATING 5 SECONDS ACTION
	fmt.Println("Worker Run function not implemented")

		// iterate over worker requests and synchronously make them

	fmt.Println(w.requests)
	return []string{w.requests[0].URL.Host+w.requests[0].URL.Path} //This function should return all founded valid urls
}

func (wp *WorkerPool) Run() []string{
	fmt.Println("Workers :", wp.workers)
	var runned = 0
    var c = make(chan bool) //Creating channel c
	go func() {
		for _, worker := range wp.workers {

			worker := worker
			go func(){
				wp.results = append(wp.results, worker.Run()...)
				runned += 1
				if runned == len(wp.workers) {
					c <- true //Sending signal in channel c if all workers done
				}
	
			}()
		}
	}()

	fmt.Println("Waiting workers...")
	<- c //Waiting signal in channel c to continue
	fmt.Println("All workers done !")

	return wp.results


}
