package main

import "fmt"

func main() {
	sliceInt := []int{6, 5, 3, 1, 8, 7, 2, 4, 3}
	sliceInt2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	insertionSort(sliceInt)
	fmt.Println(sliceInt)
	insertionSort(sliceInt2)
	fmt.Println(sliceInt2)
}

func insertionSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i] < data[j] {
				temp := data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
}
