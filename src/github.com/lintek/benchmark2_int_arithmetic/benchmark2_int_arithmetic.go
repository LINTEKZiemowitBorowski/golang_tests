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

func run() int {
	var result int = 0
	defer wg.Done()

	for x := 0; x < 1000000; x++ {
		p := x * 3
		result += ((x*x) + (p*p))/5
	}

	th_counter++

	return result
}

func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	start_time := time.Now()

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go run()
	}

	wg.Wait()
	execution_time := time.Now().Sub(start_time)
	fmt.Printf("Execution time: %f\n", execution_time.Seconds())
	fmt.Printf("Run threads: %d\n", th_counter)
}
