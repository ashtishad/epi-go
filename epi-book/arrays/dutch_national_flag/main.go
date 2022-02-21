// similar: leetcode sort colors

package epi

import "math/rand"

// Dutch national flag problem
// https://en.wikipedia.org/wiki/Dutch_national_flag_problem

// DutchFlagPartition Algorithm:
// keep the following invariants during partitioning
// Invariants:
//  - v[:low] < p
//  - v[low:mid] = p
//  - v[mid:high] are unknown
//  - v[high:] > p
//
//         < p       = p        unknown       > p
//     -----------------------------------------------
// v: |          |          |a            |           |
//     -----------------------------------------------
//                ^          ^             ^
//               low        mid           high
func DutchFlagPartition(v []int, p int) (low, high int) {
	low, high = 0, len(v)
	for mid := 0; mid < high; {

		switch a := v[mid]; {
		case a < p:
			swap(v, mid, low)
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			swap(v, mid, high-1)
			high--
		}
	}
	return
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// getPivot returns median of three random elements,pivot becomes 1.2n logn number of comparisons
func getPivot(v []int) int {
	n := len(v)
	return Median(v[rand.Intn(n)],
		v[rand.Intn(n)],
		v[rand.Intn(n)])
}

func Median(a, b, c int) int {
	return max(min(a, b), min(max(a, b), c))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
