// Optimal, time O(N), space O(1)
// Negative marking (note: num range is 1-N)

package leetcode

func findDuplicates(nums []int) []int {
	var res []int
	for _, v := range nums {
		var idx = abs(v) - 1 // range 1-N(thus, -1)
		if nums[idx] < 0 {
			res = append(res, idx+1) // range 1-N(thus, +1)
			continue
		}
		nums[idx] *= -1
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
