package main

import (
	"fmt"
	"time"
)

// these workers will receive work on the jobs channel and send the corresponding
// results on result. fibonacci calc.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// fmt.Println("worker", id, "job", j)
		fmt.Printf("Fib: %d is: %d \n", j, Fibonacci(j))
		// fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	time1 := time.Now()

	const numJobs = 40
	// Two channels to send them work and collect their results.
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// this start up 5 workers, initially blocked because there are no jobs yet
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	// here we send numJobs and the close that channel to
	// indicate that's all the work we have
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect all the results of the work
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	fmt.Println(time.Since(time1))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
