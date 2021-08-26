package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for i := range jobs {
		fmt.Println("worker", id, "started job", i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", i)
		res <- i * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, res)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-res
	}
}
