package main

import "fmt"

func main() {
	// CREATE
	m := make(map[string]float64) // empty map of [KEY] VALUES
	//fmt.Println("nil map : ", m)

	m1 := make(map[string]float64) // Empty map of string-float64 pairs
	//m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

	m1 = map[string]float64{ // Map literal
		"e":  2.71828,
		"pi": 3.1416, // notice , at the end
	}
	fmt.Println(len(m1)) // Size of map: 2

	// Add, update, get and delete keys/values

	m["pi"] = 3.14   // Add a new key-value pair
	m["pi"] = 3.1416 // Update value
	fmt.Println(m)   // Print map: "map[pi:3.1416]"

	v := m["pi"] // Get value: v == 3.1416
	v = m["pie"] // Not found: v == 0 (zero value)
	fmt.Println(v)

	// Comma OK Idiom
	fmt.Println("Comma OK Idiom starts : ")
	v, ok := m1["pi"]
	fmt.Println(v, ok) // 3.1416, true
	v, ok = m1["e"]
	fmt.Println(v, ok) // 2.71828, true
	v, ok = m1["goodbye"]
	fmt.Println(v, ok) // 0, false
	fmt.Println("Comma OK Idiom ends.")

	if x, found := m["pi"]; found {
		fmt.Println(x)
	} // Prints "3.1416"

	//delete(m, "pi") // Delete a key-value pair
	//fmt.Println(m)  // Print map: "map[]"

	// ITERATING
	for key, value := range m1 { // Order not specified
		fmt.Println(key, value)
	}
}
