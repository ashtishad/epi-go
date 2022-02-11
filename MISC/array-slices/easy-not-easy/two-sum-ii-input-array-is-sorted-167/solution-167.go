package leetcode

func twoSum(numbers []int, target int) []int {
	var n = len(numbers)
	if n == 2 {
		return []int{1, 2}
	}
	var lp, rp = 0, n - 1
	for lp < rp {
		var sum = numbers[lp] + numbers[rp]
		if sum == target {
			return []int{lp + 1, rp + 1} // idx are 1-based
		} else if sum < target {
			lp++
		} else {
			rp--
		}
	}
	return nil
}
