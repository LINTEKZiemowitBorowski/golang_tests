package main

import (
	"fmt"
	"time"
	"runtime"
	_ "github.com/howeyc/crc16"
	"github.com/howeyc/crc16"
)

const (
	ARRAY_LEN = 4
	NUM_DATA = 4
)

func prepareData() [][]uint8 {
	data := make([][]uint8, ARRAY_LEN)

	for i := 0; i < len(data); i++ {
		subItem := make([]uint8, NUM_DATA)
		for v := 0; v < len(subItem); v++ {
			subItem[v] = uint8((v + i) % 255)
		}
		data[i] = subItem
	}

	return data
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	myData := prepareData()
	fmt.Printf("myData: %v\n", myData)

	myTable := crc16.MakeTable(0xA001)

	startTime := time.Now()

	checkSums := make([]uint16, ARRAY_LEN)
	for i:=0; i < len(myData); i++ {
		checkSums[i] = crc16.Checksum(myData[i], myTable)
	}

	stopTime := time.Now()

	fmt.Printf("Execution time: %f\n", stopTime.Sub(startTime).Seconds())
	fmt.Printf("Check sums: %v\n", checkSums)
}