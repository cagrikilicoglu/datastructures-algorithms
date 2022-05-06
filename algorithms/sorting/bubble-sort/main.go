package main

import "fmt"

func main() {
	sliceInt := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sliceInt2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	bubbleSort(sliceInt)
	fmt.Println(sliceInt)
	bubbleSort(sliceInt2)
	fmt.Println(sliceInt2)
}

func bubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				temp := data[j+1]
				data[j+1] = data[j]
				data[j] = temp
			}
		}
	}
}
