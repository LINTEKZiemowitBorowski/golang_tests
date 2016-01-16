package main

import (
	"fmt"
	"os"
	"time"
	"runtime"
	"strings"
	"bytes"
)


const (
	NUM_ITERATIONS = 200000
	NUM_TASKS = 50
)


func prepareData() [] string {
	data := make([]string, 0)

    for c := 0x30; c < 0x79; c+=2 {
		item := fmt.Sprintf("%c%c%c", c, c+1, c+2)
		data = append(data, item)
	}

	return data
}

func function1(inputData []string, c chan string) {
	result := ""

	for i:= 0; i < len(inputData); i++ {
		result += inputData[i]
	}

    c <- result
	close(c)
}

func function2(inputData []string, c chan string) {
	c <- strings.Join(inputData, "")
	close(c)
}

func function3(inputData []string, c chan string) {
	var result bytes.Buffer

	for i:= 0; i < len(inputData); i++ {
		result.WriteString(inputData[i])
	}

	c <- result.String()
	close(c)
}

func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	myData := prepareData()
	/*fmt.Printf("myData: %v\n", myData)*/

	funcs := []func([]string, chan string) {function1, function2, function3}
	for f := range funcs {

		var chans [NUM_TASKS] chan string
		for i := range chans {
		   chans[i] = make(chan string)
		}

		startTime := time.Now()

		for i := range chans {
			go funcs[f](myData, chans[i])
		}

		stopTime := time.Now()

		fmt.Printf("Execution time for function%d: %f\n", f+1, stopTime.Sub(startTime).Seconds())

		/*for i := range chans {
		   fmt.Printf("Generated output for task %d: %s\n", i, <-chans[i])
		}*/
	}
}
