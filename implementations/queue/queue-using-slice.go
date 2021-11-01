package main

import (
	"fmt"
)

type Queue []int

func (q *Queue) enqueue(e int) {
	*q = append(*q, e)
}

func (q *Queue) dequeue() int {
	var old = *q
	var x = old[0] // first element
	old[0] = 0     // watch out for memory leaks, Erase element (write zero value)
	*q = old[1:]   // range: 1 to last elem
	return x
}

func (q *Queue) len() int { return len(*q) }

func main() {
	var q = &Queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	for len(*q) > 0 {
		fmt.Printf("Dequeded :  %d ", q.dequeue())
		fmt.Printf("  Len : %d", q.len())
		fmt.Println()
	}
}
