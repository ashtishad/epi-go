package algorithms

import "math/rand"

func Quicksort(nums []int) []int {
	//if len(nums) < 20 {
	//	InsertionSort(nums)
	//	return
	//}
	p := Pivot(nums)
	low, high := Partition(nums, p)
	Quicksort(nums[:low])
	Quicksort(nums[high:])
	return nums
}

// Pivot : median of three random elements,pivot becomes 1.2 nlogn number of comparisons
func Pivot(nums []int) int {
	n := len(nums)
	return Median(nums[rand.Intn(n)],
		nums[rand.Intn(n)],
		nums[rand.Intn(n)])
}

func Median(a, b, c int) int {
	var med = max(min(a, b), min(max(a, b), c))
	return med
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Partition reorders the elements of v so that:
// - all elements in v[:low] are less than p,
// - all elements in v[low:high] are equal to p,
// - all elements in v[high:] are greater than p.
func Partition(nums []int, p int) (low, high int) {
	low, high = 0, len(nums)
	for mid := 0; mid < high; {
		// Invariant:
		//  - nums[:low] < p
		//  - nums[low:mid] = p
		//  - nums[mid:high] are unknown
		//  - nums[high:] > p
		//
		//         < p       = p        unknown       > p
		//     -----------------------------------------------
		// nums: |          |          |a            |           |
		//     -----------------------------------------------
		//                ^          ^             ^
		//               low        mid           high
		switch a := nums[mid]; {
		case a < p:
			nums[mid] = nums[low]
			nums[low] = a
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			nums[mid] = nums[high-1]
			nums[high-1] = a
			high--
		}
	}
	return
}
