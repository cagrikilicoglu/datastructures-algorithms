package main

import "fmt"

func main() {

	fmt.Println(fibonacciRecursive(12))
	fmt.Println(fibonacciIterative(12))
	fmt.Println(fibonacciIterative2(12))
}

// O(n)
func fibonacciIterative2(n int) int {
	slice := []int{0, 1}
	for i := 2; i <= n; i++ {
		slice = append(slice, (slice[i-2] + slice[i-1]))
	}
	return slice[n]
}

func fibonacciIterative(n int) int {
	i := 0
	j := 1
	index := 1
	for {
		i = i + j
		index++
		if index == n {
			return i
		}
		j = j + i
		index++
		if index == n {
			return j
		}
	}
}

//O(2^n)
func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
