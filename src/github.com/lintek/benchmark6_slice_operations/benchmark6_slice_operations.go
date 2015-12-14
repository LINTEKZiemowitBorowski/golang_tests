package main

import (
	"fmt"
	"os"
	"math/rand"
	"sort"
	"time"
)

const (
	SEQUENCE_LEN = 10000
	ITERATIONS = 10
	MAX_VALUE = 10000
	GENERATOR_SEED = 1000
)


func BubbleSort(dataList []int) {
	for i := 0; i < len(dataList) - 1; i++ {
		for j := i + 1; j < len(dataList); j++ {
			if dataList[i] > dataList[j] {
				dataList[i], dataList[j] = dataList[j], dataList[i]
			}
		}
	}
}


func BuiltInSort(dataList []int) {
	sort.Ints(dataList)
}


func FillList() []int {
	generator := rand.New(rand.NewSource(GENERATOR_SEED))
	subList := make([]int, SEQUENCE_LEN)
	for y := 0; y < len(subList); y++ {
		subList[y] = generator.Intn(MAX_VALUE)
	}

	return subList
}


func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	var srcList = make([][]int, ITERATIONS)
	var tmpList = make([][]int, ITERATIONS)

	// Create test data
	for x := 0; x < len(srcList); x++ {
		srcList[x] = FillList()
	}

	for i := 0; i < 2; i++ {

		// fmt.Printf("Unsorted data: %v\n", srcList)
		startTime := time.Now()

		for x := 0; x < len(tmpList); x++ {
			tmpList[x] = make([]int, SEQUENCE_LEN)
			copy(tmpList[x], srcList[x])
		}

		sortTime := time.Now()

		if i == 0 {
			for _, subList := range tmpList {
				BuiltInSort(subList)
			}
		} else {
			for _, subList := range tmpList {
				BubbleSort(subList)
			}
		}

		endTime := time.Now()

		copyingTime := sortTime.Sub(startTime)
		sortingTime := endTime.Sub(sortTime)

		fmt.Printf("Copying time: %f\n", copyingTime.Seconds())
		fmt.Printf("Sorting time: %f\n", sortingTime.Seconds())
		// fmt.Printf("Sorted data: %v\n", tmpList)
	}
}