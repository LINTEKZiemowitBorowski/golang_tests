package main

import (
	"fmt"
	"time"
	"sort"
)

var myList = [][]int{
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20},
	{54, 26, 93, 17, 77, 31, 44, 55, 20}}


func BuiltInSort(dataList []int) {
	sort.Ints(dataList)
}


func main() {
	var tmpList = make([][][]int, 10000)

	// fmt.Printf("Unsorted data: %v\n", myList)
	startTime := time.Now()

	for x := 0; x < len(tmpList); x++ {
		tmpList[x] = make([][]int, len(myList) )
		for y := 0; y < len(myList); y++ {
			tmpList[x][y] = make([]int, len(myList[y]))
			copy(tmpList[x][y], myList[y])
		}
	}

	sortTime := time.Now()

	for _, listCopy := range tmpList {
		for _, subList := range listCopy {
				BuiltInSort(subList)
			}
	}

	endTime := time.Now()

	copyingTime := sortTime.Sub(startTime)
	sortingTime := endTime.Sub(sortTime)

	fmt.Printf("Copying time: %f\n", copyingTime.Seconds())
	fmt.Printf("Sorting time: %f\n", sortingTime.Seconds())
	// fmt.Printf("Sorted data: %v\n", tmpList)
}
