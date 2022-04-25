package main

import (
	"fmt"
	"log"
)

func main() {
	// testSlice1 := []int{1, 2, 3, 9}
	// testSlice2 := []int{1, 2, 4, 4}

	testSlice3NotSorted := []int{10, 2, 3, 1, 9}
	testSlice4NotSorted := []int{3, 2, 1, 5}
	// findSumOfPairsSorted(testSlice1, 8)
	// findSumOfPairsSorted(testSlice2, 8)
	// findSumOfPairsSorted2(testSlice1, 8)
	// findSumOfPairsSorted2(testSlice2, 8)
	result1 := findSumOfPairsNotSorted(testSlice3NotSorted, 8)
	result2 := findSumOfPairsNotSorted(testSlice4NotSorted, 8)
	fmt.Println(result1, result2)
}

func findSumOfPairsNotSorted(slice []int, sum int) bool {
	var remainder int
	remainderMap := make(map[int]bool)
	for i := range slice {

		if remainderMap[slice[i]] {
			return true
		}
		remainder = sum - slice[i]
		remainderMap[remainder] = true

	}
	return false

}

// Better solution if sorted
func findSumOfPairsSorted2(slice []int, output int) {
	i, j := 0, 0
	for {
		if i >= len(slice) || j >= len(slice) {
			break
		}
		if slice[i]+slice[(len(slice)-1)-j] > output {
			j++
		} else if slice[i]+slice[(len(slice)-1)-j] < output {
			i++
		} else {
			log.Println("Solution found")
			log.Println(i, (len(slice)-1)-j)
			return
		}

	}
	log.Println("Solution not found")
}

func findSumOfPairsSorted(slice []int, output int) {
	for i := range slice {
		for j := i + 1; j < len(slice); j++ {
			if slice[i]+slice[j] == output {
				log.Println(i, j)
				log.Println("Solution found")
				return
			}
		}
	}
	log.Println("Solution not found")

}
