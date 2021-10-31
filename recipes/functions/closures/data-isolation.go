// Golang program to illustrate how
// to create data isolation
package main

import "fmt"

// newCounter function to
// isolate global variable
func newCounter() func() int {
	GFG := 0
	return func() int {
		GFG += 1
		return GFG
	}
}
func main() {
	// newCounter function is
	// assigned to a variable
	counter := newCounter()

	// invoke counter
	fmt.Println(counter()) // 1
	// invoke counter
	fmt.Println(counter()) // 2
}
