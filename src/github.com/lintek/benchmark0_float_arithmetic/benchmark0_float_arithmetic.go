package main

import (
	"fmt"
	"os"
	"time"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

var th_counter int = 0

func run() float32 {
	var result float32 = 0
	defer wg.Done()

	for x := 0; x < 1000000; x++ {
		var f float32 = float32(x)
		var p float32 = f * 3.14
		result += ((f * f) + (p * p)) / 5.5
	}

	th_counter++

	return result
}

func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	wg.Add(50)
	start_time := time.Now()

	for i := 0; i < 50; i++ {
		go run()
	}

	wg.Wait()
	execution_time := time.Now().Sub(start_time)
	fmt.Printf("Execution time: %f\n", execution_time.Seconds())
	fmt.Printf("Run threads: %d\n", th_counter)
}
