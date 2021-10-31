package main

import "fmt"

func main() {
	type Student struct {
		Name string
		Age  int
	}

	var ash Student                // ash == Student{"", 0}
	ash.Name = "Alice"             // ash == Student{"Alice", 0}
	fmt.Println(ash.Name, ash.Age) // Alice, 0

	// 2 ways to create and initialize a new struct

	// WAY:1 The new keyword can be used to create a new struct.
	//It returns a pointer to the newly created struct.
	var pa *Student   // pa == nil
	pa = new(Student) // pa == &Student{"", 0}
	pa.Name = "Alice" // pa == &Student{"Alice", 0}

	// WAY:2 You can also create and initialize a struct with a struct literal.
	b := Student{ // b == Student{"Bob", 0}
		Name: "Bob",
	}

	pb := &Student{ // pb == &Student{"Bob", 8}
		Name: "Bob",
		Age:  8,
	}

	c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
	d := Student{}             // d == Student{"", 0}

	fmt.Println(b, pb, c, d)

	// Compare structs

	d1 := Student{"David", 1}
	d2 := Student{"David", 2}
	fmt.Println(d1 == d2) // false
}
