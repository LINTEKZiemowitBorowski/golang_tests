package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Create random generator
var generator = rand.New(rand.NewSource(1000))


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
	subList := make([]int, 500)
	for y := 0; y < len(subList); y++ {
		subList[y] = generator.Int()
	}

	return subList
}


func main() {
	var srcList = make([][]int, 5)
	var tmpList = make([][][]int, 10000)

	// Create test data
	for x := 0; x < len(srcList); x++ {
		srcList[x] = FillList()
	}

	for i := 0; i < 2; i++ {

		// fmt.Printf("Unsorted data: %v\n", srcList)
		startTime := time.Now()

		for x := 0; x < len(tmpList); x++ {
			tmpList[x] = make([][]int, len(srcList))
			for y := 0; y < len(srcList); y++ {
				tmpList[x][y] = make([]int, len(srcList[y]))
				copy(tmpList[x][y], srcList[y])
			}
		}

		sortTime := time.Now()

		if i == 0 {
			for _, listCopy := range tmpList {
				for _, subList := range listCopy {
					BuiltInSort(subList)
				}
			}
		} else {
			for _, listCopy := range tmpList {
				for _, subList := range listCopy {
					BubbleSort(subList)
				}
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

