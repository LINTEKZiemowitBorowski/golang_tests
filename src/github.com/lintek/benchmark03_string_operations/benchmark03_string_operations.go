package main

import (
	"fmt"
	"os"
	"time"
	"runtime"
	"sync"
	"strings"
	"bytes"
)

var wg sync.WaitGroup

const (
	NUM_ITERATIONS = 100000
	NUM_TASKS = 50
)


func prepareData() [] string {
	data := make([]string, 0)

    for c := 0x30; c < 0x79; c+=3 {
		item := fmt.Sprintf("%c%c%c", c, c+1 , c+2)
		data = append(data, item)
	}

	return data
}

func function1(inputData []string) {
	result := ""
	defer wg.Done()

	for i:= 0; i < len(inputData); i++ {
		result += inputData[i]
	}
}

func function2(inputData []string) {
	defer wg.Done()
	strings.Join(inputData, "")
}

func function3(inputData []string) {
	var result bytes.Buffer
	defer wg.Done()

	for i:= 0; i < len(inputData); i++ {
		result.WriteString(inputData[i])
	}

	result.String()
}

func main() {
	funcs := []func([]string) {function1, function2, function3}
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	myData := prepareData()
	fmt.Printf("myData: %v\n", myData)

	for f := 0; f < 3; f++ {
		wg.Add(50)
		startTime := time.Now()

		for i := 0; i < NUM_TASKS; i++ {
			go funcs[f](myData)
		}

		wg.Wait()
		stopTime := time.Now()

		fmt.Printf("Execution time for function%d: %f\n", f+1, stopTime.Sub(startTime).Seconds())
	}

}
