package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0, // NOTICE this COMMA at the end
	}
	v, ok := m["hello"]
	fmt.Println(v, ok) // 5, true
	v, ok = m["world"]
	fmt.Println(v, ok) // 0, true
	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0, false

}
