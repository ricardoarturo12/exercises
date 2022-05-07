package main

import "fmt"

type Job struct {
	valueA int
	valueB int
}

type Result struct {
	job       Job
	sumResult int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		output := Result{job: job, sumResult: mult(job.valueA, job.valueB)}
		// fmt.Printf("Calc id: %d valueA %d , valueB %d, Result: %d\n", id, job.valueA, job.valueB, output.sumResult)
		results <- output
	}
}

func main() {
	numWorker := 10
	jobs := make(chan Job, numWorker)
	results := make(chan Result, numWorker)

	for i := 1; i <= numWorker; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i <= 5; i++ {
		jobs <- Job{valueA: i, valueB: i + 1}
	}
	close(jobs)

	for i := 0; i <= 5; i++ {
		// <-results
		fmt.Println(<-results)
	}

}

func mult(a, b int) int {
	return a * b
}
