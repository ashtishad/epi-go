package ds

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	// put h into old[]int -> len old -> put last elem to x -> restore old to *h -> return x
	var old = *h
	var n = len(old)
	var x = old[n-1]  // last element
	*h = old[0 : n-1] // range: 0 to 2nd last elem
	return x
}

//func main() {
//	var minHeap = &IntHeap{4, 5, 6, 1, 7, 8, 2, 9}
//	heap.Init(minHeap)
//	heap.Push(minHeap, 3)
//	fmt.Println("minimum", (*minHeap)[0])
//	for minHeap.len() > 0 {
//		fmt.Printf(" %d ", heap.Pop(minHeap))
//	}
//}
