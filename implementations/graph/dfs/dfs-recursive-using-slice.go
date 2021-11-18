package main

import "fmt"

type Vertex struct {
	visited    bool
	value      int
	neighbours []*Vertex
}

func NewVertex(value int) *Vertex {
	return &Vertex{
		value: value,
		// below lines can be deleted, because this will be initialized with their zero value
		visited:    false,
		neighbours: nil, // comment 5.
	}
}

func (v *Vertex) connect(vertex ...*Vertex) { // see comment 4.
	v.neighbours = append(v.neighbours, vertex...)
}

type Graph struct{}

func (g *Graph) dfs(vertex *Vertex) {
	if vertex.visited { // see comment 1.
		return // see comment 2.
	}
	vertex.visited = true
	fmt.Println(vertex.value, vertex.visited, vertex.neighbours)
	for _, v := range vertex.neighbours { // see comment 3.
		g.dfs(v)
	}
}

func (g *Graph) disconnected(vertices ...*Vertex) {
	for _, v := range vertices {
		g.dfs(v)
	}
}

// func main() {
// 	v1 := NewVertex(1)
// 	v2 := NewVertex(2)
// 	v3 := NewVertex(3)
// 	v4 := NewVertex(4)
// 	v5 := NewVertex(5)
// 	v6 := NewVertex(6)
// 	v7 := NewVertex(7)
// 	v8 := NewVertex(8)
// 	v9 := NewVertex(9)
// 	v10 := NewVertex(10)
// 	g := Graph{}
// 	v1.connect(v2, v3, v4)
// 	v2.connect(v5, v6)  // see comment 4.
// 	v4.connect(v7, v8)  // see comment 4.
// 	v5.connect(v9, v10) // see comment 4.
// 	g.dfs(v1)
// }

/* Comments:
1. if vertex.visited != true { is the same as if !vertex.visited {
2. first handle specific cases with an early return and then the general case: this way augmented indentation concerns only the specific cases
3. you can iterate over an empty (or even nil) slice
4. use unpacking
5. I don't think that initializing the slice capacity brings much
*/
