package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

var values_result = make(map[int]int)

func Fibonacci(n int) int {
	if n <= 1 {
		values_result[n] = n
	}
	if val, ok := values_result[n]; ok {
		return val
	}
	values_result[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return values_result[n]
}

func main() {
	time1 := time.Now()
	// position 0 es la ubicación archivo
	input_value, _ := strconv.Atoi(os.Args[1])
	fib := Fibonacci(input_value)
	fmt.Printf("Fib: %d, is: %d\n", input_value, fib)
	defer fmt.Print(time.Since(time1))

	var keys []int

	// imprime solo el key -> y agrega a un slice
	for k := range values_result {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// recorre el slice e imprime el resultado
	for k := range keys {
		fmt.Printf("Fib: %d, value: %d\n", k, values_result[k])
	}
}
