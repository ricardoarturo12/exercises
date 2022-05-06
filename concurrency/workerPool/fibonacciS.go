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

func main(){
	time1 := time.Now()
	values := []int{2, 3, 5, 7, 11, 13, 40, 20, 30, 35,45}
	for _, value := range values {
		fmt.Println(Fibonacci(value))
	}

	fmt.Print(time.Since(time1))
}