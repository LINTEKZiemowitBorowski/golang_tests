package main

import (
	"fmt"
	"os"
	"time"
	"runtime"
	"github.com/jiguorui/crc16"
)

const (
	ARRAY_LEN = 255
	NUM_ITEMS = 10000
)

func prepareData() [][]uint8 {
	data := make([][]uint8, NUM_ITEMS)

	for i := range(data) {
		subItem := make([]uint8, ARRAY_LEN)
		for v := range(subItem) {
			subItem[v] = uint8((v + i) % 255)
		}
		data[i] = subItem
	}

	return data
}

func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	myData := prepareData()
	// fmt.Printf("myData: %v\n", myData)

	startTime := time.Now()

	checkSums := make([]string, NUM_ITEMS)
	for i := range(myData){
		checkSums[i] = fmt.Sprintf("%X", crc16.CheckSum(myData[i]))
	}

	stopTime := time.Now()

	fmt.Printf("Execution time: %f\n", stopTime.Sub(startTime).Seconds())
	// fmt.Printf("Check sums: %v\n", checkSums)
}