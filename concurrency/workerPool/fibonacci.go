package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	id      int
	integer int
}

type Result struct {
	job   Job
	total int
}

var size = runtime.GOMAXPROCS(0)
var jobs = make(chan Job, size)
var results = make(chan Result, size)

func worker(wg *sync.WaitGroup) {
	for c := range jobs {
		fib := Fibonacci(c.integer)
		results <- Result{c, fib}
		fmt.Printf("realizando fib: %d resultado: %d\n", c.integer, fib)
	}
	wg.Done()
}

func create(vals []int) {
	for i, value := range vals {
		c := Job{i, value}
		fmt.Printf("worker: %d, value: %d\n", i, value)
		jobs <- c
	}
	close(jobs)
}

func main() {
	time1 := time.Now()

	values := []int{ 44, 45, 46, 47, 48, 49, 50}

	nWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Task crea las tareas -> mete los valores a procesar en client  y eso luego se envía a jobs
	go create(values)

	// Pool
	finished := make(chan interface{})
	go func() {
		for d := range results {
			fmt.Printf("client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tFib: %d\n", d.job.integer, d.total)
			// fmt.Println(d)
		}
		finished <- true
	}()

	// Worker -> realiza la tarea
	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
	fmt.Printf("finished: %v\n", <-finished)
	fmt.Println(time.Since(time1))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
