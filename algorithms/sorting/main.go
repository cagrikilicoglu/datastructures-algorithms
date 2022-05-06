package main

import (
	"fmt"
	"sort"
)

func main() {

	sliceStr := []string{"a", "d", "z", "e", "r", "b"}
	sliceInt := []int{2, 65, 34, 2, 1, 7, 8}
	sort.Strings(sliceStr)
	sort.Ints(sliceInt)
	fmt.Println(sliceStr)
	fmt.Println(sliceInt)
}
