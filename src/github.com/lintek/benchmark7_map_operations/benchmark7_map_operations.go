package main

import (
	"fmt"
	"time"
	"math/rand"
)


const (
	SEQUENCE_LEN = 1000000
	MAX_VALUE = 9999999
	GENERATOR_SEED = 99999999
)


func GetRandoms() []int {
	generator := rand.New(rand.NewSource(GENERATOR_SEED))
	randoms := make([]int, SEQUENCE_LEN)
	for y := 0; y < len(randoms); y++ {
		randoms[y] = generator.Intn(MAX_VALUE)
	}

	return randoms
}


func buildMap(myKeys []int) map[int]string {
	myMap := make(map[int]string)

	for _, key := range(myKeys) {
		myMap[key] =  fmt.Sprintf("%d", key)
	}

	return myMap
}


func searchMap(myKeys []int, myMap map[int]string) []string {
	foundValues := make([]string, 0)

	for _, key := range(myKeys) {
		foundValues = append(foundValues, myMap[key])
	}

	return foundValues
}


func main() {

	createdKeys := GetRandoms()

	fmt.Printf("createdKeys len: %d\n", len(createdKeys))
	// fmt.Printf("createdKeys: %v\n", createdKeys)

	startTime := time.Now()
	randomMap := buildMap(createdKeys)

	fmt.Printf("Map len: %d\n", len(randomMap))

	buildTime := time.Now()
	foundValues := searchMap(createdKeys, randomMap)
	searchTime := time.Now()

	fmt.Printf("Found values len: %d\n", len(foundValues))

	if len(createdKeys) != len(foundValues) {
		fmt.Printf("Error, len of createdKeys(%d) is different then len of keys found in the map (%d)\n",
		len(createdKeys), len(foundValues))
	}

	fmt.Printf("Build map time: %f\n", buildTime.Sub(startTime).Seconds())
	fmt.Printf("Search map time: %f\n", searchTime.Sub(startTime).Seconds())
}
