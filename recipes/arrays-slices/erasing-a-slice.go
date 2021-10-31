package main

import "fmt"

func main() {
	// This will release the underlying array to the garbage collector.
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a)) // [] 0 0

	// #Keep allocated memory
	b := []string{"A", "B", "C", "D", "E"}
	b = b[:0]
	fmt.Println(b, len(b), cap(b)) // [] 0 5

	// If the slice is extended again, the original data reappears.
	fmt.Println(a[:2]) // [A B]

	// nil slice vs empty slice
	var ns []int = nil
	var es []int = make([]int, 0)

	fmt.Println(ns == nil) // true
	fmt.Println(es == nil) // false

	fmt.Printf("%#v\n", ns) // []int(nil)
	fmt.Printf("%#v\n", es) // []int{}

	// from go official wiki
	/* The official Go wiki recommends using nil slices over empty slices.

	[…] the nil slice is the preferred style.

	Note that there are limited circumstances where a non-nil but zero-length slice is
	preferred, such as when encoding JSON objects (a nil slice encodes to null,
	while []string{} encodes to the JSON array []).

	When designing interfaces, avoid making a distinction between a nil slice
	and a non-nil, zero-length slice, as this can lead to subtle programming errors.

	Link: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices

	*/
}
