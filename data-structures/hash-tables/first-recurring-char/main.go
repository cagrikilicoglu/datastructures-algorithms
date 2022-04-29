package main

import (
	"fmt"
)

func main() {
	slice1 := []interface{}{2, 5, 1, 2, 3, 1, 2, 4}
	slice2 := []interface{}{2, 1, 1, 2, 3, 3, 5}
	slice3 := []interface{}{1, 2, 3}
	slice4 := []interface{}{}
	fmt.Println(findFirstRecurring(slice1))
	fmt.Println(findFirstRecurring(slice2))
	fmt.Println(findFirstRecurring(slice3))
	fmt.Println(findFirstRecurring(slice4))

	fmt.Println(findFirstRecurringBrute(slice1))
	fmt.Println(findFirstRecurringBrute(slice2))
	fmt.Println(findFirstRecurringBrute(slice3))
	fmt.Println(findFirstRecurringBrute(slice4))
}

func findFirstRecurring(data []interface{}) interface{} {

	sliceMap := make(map[interface{}]bool)

	for i := range data {
		if sliceMap[data[i]] {
			// fmt.Printf("First recurring character is %v, at index %d", data[i], i)
			return data[i]

		} else {
			sliceMap[data[i]] = true
		}
	}
	return nil
}

func findFirstRecurringBrute(data []interface{}) interface{} {

	minj := len(data)
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[i] == data[j] {
				if minj > j {
					minj = j
				}
			}
		}
	}
	if minj != 0 && minj < len(data) {
		return data[minj]
	} else {
		return nil
	}
}
