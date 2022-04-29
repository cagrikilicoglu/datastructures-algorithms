package main

import "fmt"

func main() {

	slice1 := []interface{}{"x", "y", "z", 1}
	slice2 := []interface{}{"a", "b", 1}

	// result := checkCommonElement(slice1, slice2)
	result := checkCommonElement2(slice1, slice2)
	fmt.Println(result)
}

// better solution O(a+b)
func checkCommonElement2(arr1, arr2 []interface{}) bool {
	arrMap := make(map[interface{}]bool)
	for i := range arr1 {
		if !arrMap[arr1[i]] {
			arrMap[arr1[i]] = true
		}
	}

	for i := range arr2 {
		if arrMap[arr2[i]] {
			return true
		}
	}
	return false

}

// O(n^2)
func checkCommonElement(arr1, arr2 []interface{}) bool {

	for i := range arr1 {
		for j := range arr2 {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}
