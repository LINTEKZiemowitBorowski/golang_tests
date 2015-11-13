package main

import (
	"fmt"
	"math"
	"time"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

var th_counter int = 0

func run() float64 {
	var result float64 = 0
	defer wg.Done()

	for x := 0; x < 1000000; x++ {
		var f float64 = float64(x)
		var p float64 = f * math.Pi
		result += math.Sqrt(math.Pow(f,2) + math.Pow(p,2))
	}

	th_counter++

	return result
}

func main() {

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
