package main

import "fmt"

func main() {

	// a complete for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// a condition only for
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// for range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

}
