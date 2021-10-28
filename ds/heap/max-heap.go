package ds

type IntMaxHeap []int

func (h IntMaxHeap) Len() int { return len(h) }

// Less just change in less function
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	// put h into old[]int -> len old -> put last elem to x -> restore old to *h -> return x
	var old = *h
	var n = len(old)
	var x = old[n-1]  // last element
	*h = old[0 : n-1] // range: 0 to 2nd last elem
	return x
}

//func main() {
//	var maxHeap = &IntMaxHeap{4, 5, 6, 1, 7, 8, 2, 9}
//	heap.Init(maxHeap)
//	heap.Push(maxHeap, 3)
//	fmt.Println("minimum", (*maxHeap)[0])
//	for maxHeap.Len() > 0 {
//		fmt.Printf(" %d ", heap.Pop(maxHeap))
//	}
//}
