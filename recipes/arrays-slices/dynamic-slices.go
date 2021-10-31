package main

import "fmt"

func main() {
	var size = 10
	var arInt = make([]int, size)
	arInt[0] = 12

	fmt.Println(arInt)
	// [12 0 0 0 0 0 0 0 0 0]
}
