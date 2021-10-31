package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Robbert", 18},
		{"Fred", "Freon", 18},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	// SliceStable ensures same value won't change order
	sort.SliceStable(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
