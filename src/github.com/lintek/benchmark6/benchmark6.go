package main

import (
	"fmt"
	"time"
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


func BubbleSort(dataList []int) {
	for i := 0; i < len(dataList) - 1; i++ {
		for j := i + 1; j < len(dataList); j++ {
			if dataList[i] > dataList[j] {
				dataList[i], dataList[j] = dataList[j], dataList[i]
			}
		}
	}
}


func main() {
	var tmpList = make([][][]int, 10000)

	fmt.Printf("Unsorted data: %v\n", myList)
	startTime := time.Now()


	for x := 0; x < 10000; x++ {
		tmpList[x] = make([][]int, len(myList) * 9)
		copy(tmpList[x], myList)
	}

	sortTime := time.Now()

	for x := 0; x < 10000; x++ {
		for _, element := range tmpList[x] {
				BubbleSort(element)
			}
	}

	endTime := time.Now()

	copyingTime := sortTime.Sub(startTime)
	sortingTime := endTime.Sub(sortTime)

	fmt.Printf("Copying time: %f\n", copyingTime.Seconds())
	fmt.Printf("Sorting time: %f\n", sortingTime.Seconds())
	fmt.Printf("Sorted data: %v\n", myList)
}
