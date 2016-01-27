package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
)


const (
	SEQUENCE_LEN = 1000000
	MAX_VALUE = 99999999
)


func GetRandoms() []int64 {
	generator := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	randoms := make([]int64, SEQUENCE_LEN)
	for y := range(randoms) {
		randoms[y] = generator.Int63n(MAX_VALUE)
	}

	return randoms
}


func buildMap(myKeys []int64) map[int64]string {
	myMap := make(map[int64]string)

	for _, key := range(myKeys) {
		myMap[key] =  fmt.Sprintf("%d", key)
	}

	return myMap
}


func searchMap(myKeys []int64, myMap map[int64]string) []string {
	myValues := make([]string, 0)

	for _, key := range(myKeys) {
		myValues = append(myValues, myMap[key])
	}

	return myValues
}


func main() {
	fmt.Printf("Running: %s\n", os.Args[0])

	createdKeys := GetRandoms()

	// fmt.Printf("createdKeys len: %d\n", len(createdKeys))
	// fmt.Printf("createdKeys: %v\n", createdKeys)

	startTime := time.Now()
	randomMap := buildMap(createdKeys)

	// fmt.Printf("Map len: %v\n", len(randomMap))

	buildTime := time.Now()
	foundValues := searchMap(createdKeys, randomMap)
	searchTime := time.Now()

	// fmt.Printf("Found values len: %d\n", len(foundValues))

	if len(createdKeys) != len(foundValues) {
		fmt.Printf("Error, len of createdKeys(%d) is different then len of keys found in the map (%d)\n",
		len(createdKeys), len(foundValues))
	}

	fmt.Printf("Build map time: %f\n", buildTime.Sub(startTime).Seconds())
	fmt.Printf("Search map time: %f\n", searchTime.Sub(startTime).Seconds())
}
