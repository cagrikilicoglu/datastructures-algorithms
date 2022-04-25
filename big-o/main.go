package main

import (
	"log"
	"time"
)

func main() {

	// nemoHun := make([]string, 100000)
	// nemo := []string{"nemo"}
	// fillSlice(nemoHun, "nemo")
	// findNemo(nemoHun) // O(n) -> Linear Time

	boxes := []int{0, 1, 2, 3, 4, 5, 6}
	// printFirst(boxes)    // O(1) -> Constant Time
	// printFirstTwo(boxes) // O(2) -> Constant Time
	logAllPairs(boxes)
}

func logAllPairs(arr []int) {
	for i := range arr {
		// for j := range arr {
		// 	log.Println(arr[i], arr[j])
		// }
		log.Printf("%d", arr[i])
	}

}

func printFirst(arr []int) {
	log.Println(arr[0])
}
func printFirstTwo(arr []int) {
	log.Println(arr[0], arr[1])
}

func findNemo(arr []string) {
	t1 := time.Now()
	for i := range arr {
		if arr[i] == "nemo" {
			log.Println("Found Nemo!")
		}
	}
	t2 := time.Now()

	performance := t2.Sub(t1)
	log.Println(performance)
}

func fillSlice(arr []string, fill string) {
	for i := range arr {
		arr[i] = fill
	}
}
