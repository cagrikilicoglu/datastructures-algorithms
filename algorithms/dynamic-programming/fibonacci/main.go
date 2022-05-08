package main

import "log"

func main() {
	fasterFibo := fibonacciMaster()
	log.Println(fasterFibo(8))
	log.Println(fibonacci(8))
	log.Println(calculations)
	log.Println(cachceFibonacci)
}

var cachceFibonacci = make(map[int]int)
var calculations int

func fibonacci(index int) int {

	if index < 2 {
		return index
	}
	return fibonacci(index-1) + fibonacci(index-2)
}
func fibonacciMaster() func(index int) int {
	cacheFibo := make(map[int]int)
	var fib func(int) int
	fib = func(index int) int {
		calculations++

		if _, ok := cacheFibo[index]; ok {
			return cacheFibo[index]
		} else {
			if index < 2 {
				return index
			} else {
				cacheFibo[index] = fib(index-1) + fib(index-2)
				return cacheFibo[index]
			}

		}

	}
	return fib
}
