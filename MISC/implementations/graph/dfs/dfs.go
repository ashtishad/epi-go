package main

import "fmt"

type Node struct {
	Value    int
	adjacent []*Node
}

func (n *Node) DepthFirstSearch(g []int) []int {
	g = append(g, n.Value)
	for _, adj := range n.adjacent {
		g = adj.DepthFirstSearch(g)
	}
	return g
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func main() {
	// DFS: [1,2,5,9,10,6,3,4,7,8]
	v1 := NewNode(1)
	v2 := NewNode(2)
	v3 := NewNode(3)
	v4 := NewNode(4)
	v5 := NewNode(5)
	v6 := NewNode(6)
	v7 := NewNode(7)
	v8 := NewNode(8)
	v9 := NewNode(9)
	v10 := NewNode(10)
	v1.adjacent = []*Node{v2, v3, v4}
	v2.adjacent = []*Node{v5, v6}
	v4.adjacent = []*Node{v7, v8}
	v5.adjacent = []*Node{v9, v10}
	var in = make([]int, 0, 10)
	var res = v1.DepthFirstSearch(in)
	fmt.Println(res)
}
