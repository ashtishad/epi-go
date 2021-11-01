package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) push(e int) {
	*s = append(*s, e)
}

func (s *Stack) pop() int {
	var old = *s
	var n = len(old)
	var x = old[n-1]  // last element
	old[n-1] = 0      // watch out for memory leaks, Erase element (write zero value)
	*s = old[0 : n-1] // range: 0 to 2nd last elem
	return x
}

func (s *Stack) len() int { return len(*s) }

func main() {
	var s = &Stack{}
	s.push(1)
	s.push(2)
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	for len(*s) > 0 {
		fmt.Printf("Popped :  %d ", s.pop())
		fmt.Printf("  Len : %d", s.len())
		fmt.Println()
	}
}
