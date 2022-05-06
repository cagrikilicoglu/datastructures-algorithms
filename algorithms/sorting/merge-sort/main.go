package main

import "fmt"

func main() {
	sliceInt := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sliceInt2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	fmt.Println(mergeSort(sliceInt))

	fmt.Println(mergeSort(sliceInt2))
}

func mergeSort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}
	var midIndex int
	if len(slice)%2 == 1 {
		midIndex = (len(slice) + 1) / 2
	} else {
		midIndex = (len(slice) / 2)
	}

	left := slice[:midIndex]
	right := slice[midIndex:]

	return merge(mergeSort(left), mergeSort(right))

}

func merge(left, right []int) []int {
	result := []int{}

	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {

			result = append(result, left[leftIndex])
			leftIndex++

		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex == len(left) {
		result = append(result, right[rightIndex:]...)
	} else if rightIndex == len(right) {
		result = append(result, left[leftIndex:]...)
	}
	return result
}
