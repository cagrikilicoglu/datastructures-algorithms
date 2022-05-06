package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello"))
}

func reverseString(s string) string {

	if s == "" {
		return ""
	}
	return reverseString(s[1:]) + string(s[0])
}
