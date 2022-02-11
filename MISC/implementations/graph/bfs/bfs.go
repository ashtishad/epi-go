// Time : O(V + E)

package main

import "fmt"

func bfs(start int, adjs map[int][]int) []int {
	seen := map[int]bool{}
	frontier := []int{start}
	//parent := map[int]int{}
	var res []int

	for 0 < len(frontier) {
		var nextf []int
		for _, vertex := range frontier {
			seen[vertex] = true
			res = append(res, vertex)
			//frontiers(vertex)
			var frontiers = bfsFrontier(vertex, adjs, seen)
			for _, fv := range frontiers {
				//if _, ok := parent[fv]; !ok {
				//	parent[fv] = vertex
				//}
				nextf = append(nextf, fv)
			}
		}
		frontier = nextf
	}
	return res
}

func bfsFrontier(vertex int, adjs map[int][]int, seen map[int]bool) []int {
	var nextf []int
	iter := func(n int) bool { _, ok := seen[n]; return !ok }
	for _, n := range adjs[vertex] {
		if iter(n) {
			nextf = append(nextf, n)
		}
	}
	return nextf
}

func main() {
	// BFS = {1,(2,3,4),(5,6,,7,8),(9,10)
	var adjs = map[int][]int{
		1: []int{5, 2, 3},
		2: []int{6, 7},
		7: []int{12, 13},
		5: []int{9, 11},
	}
	var r = bfs(1, adjs)
	fmt.Println(r)
}
