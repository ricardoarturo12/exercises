package main

import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	time1 := time.Now()
	// values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9 , 10, 13, 40, 20, 30, 35,45}
	for i := 1; i <= 40; i++ {
		fmt.Printf("Fib: %d, is: %d\n", i, Fibonacci(i))
	}

	fmt.Print(time.Since(time1))
}
