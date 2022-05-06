package main

import "fmt"

func main() {
	sliceInt := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sliceInt2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	selectionSort(sliceInt)
	fmt.Println(sliceInt)
	selectionSort(sliceInt2)
	fmt.Println(sliceInt2)
}

func selectionSort(data []int) {
	// min := data[0]
	var indexMin int
	for i := 0; i < len(data)-1; i++ {
		min := data[i]
		for j := i + 1; j < len(data); j++ {
			if min > data[j] {
				min = data[j]
				indexMin = j
			}
		}
		data[indexMin] = data[i]
		data[i] = min
	}
}
