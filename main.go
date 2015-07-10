package main

import (
	"log"
	"time"
)

func main() {
	jobcount := 100
	workercount := 10
	jobs := make(chan int)
	results := make(chan int)

	// Send a bunch of jobs over
	log.Println("Scheduling insertions into jobs")
	go func() {
		for i := 0; i < jobcount; i++ {
			log.Printf("Inserting %v into jobs\n", i)
			jobs <- i
		}
	}()

	for i := 1; i <= workercount; i++ {
		log.Printf("Scheduling worker %v\n", i)
		go worker(i, jobs, results)
	}

	for i := 0; i < jobcount; i++ {
		log.Printf("Yielding results %v\n", i)
		<-results
	}
}

func worker(id int, jobs, results chan int) {
	for jobid := range jobs {
		log.Printf("worker %v : job %v\n", id, jobid)
		time.Sleep(1 * time.Second)

		results <- jobid
	}
}
