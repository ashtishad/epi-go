package main

import "fmt"

func main() {
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3, 10}
	fmt.Println(addTo(5, a...))                    // [9 8 15]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
}

// The variadic parameter must be the last (or only) parameter
// in the input parameter list. E.g: println, append function
func addTo(base int, values ...int) []int { // variadic parameter
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, base+v)
	}
	return out
}
