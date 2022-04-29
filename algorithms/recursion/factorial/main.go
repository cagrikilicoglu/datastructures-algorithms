package main

import "fmt"

func main() {

	fmt.Println(findFactorialIterative(5))
	fmt.Println(findFactorialRecursive(0))
}

// O(n)
func findFactorialRecursive(num int) int {

	var result int
	if num <= 1 {
		return 1
	}

	result = (num) * findFactorialRecursive(num-1)
	return result

}

// O(n)
func findFactorialIterative(num int) int {

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
