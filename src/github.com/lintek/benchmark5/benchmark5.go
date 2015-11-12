package main

import (
	"time"
	"fmt"
	"runtime"
)

func fibonacci(n int) int {
	if (n <= 1) {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main () {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	start_time := time.Now()
	fibonacci(40)
	execution_time := time.Now().Sub(start_time)
	fmt.Printf("Execution time: %f\n", execution_time.Seconds())
}
