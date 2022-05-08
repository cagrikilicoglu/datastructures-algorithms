package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(addTo80(5))
	// fmt.Println(addTo80(5))

	memoize := memoizedAddTo80()
	fmt.Println(memoize(5))
	fmt.Println(memoize(5))
}

func addTo80(n int) int {
	log.Println("long time")
	return n + 80
}

func memoizedAddTo80() func(n int) int {
	var cache = make(map[int]int)
	return func(n int) int {
		if _, ok := cache[n]; ok {
			return cache[n]
		} else {
			log.Println("long time")
			cache[n] = 80 + n
			return cache[n]
		}

	}
}
