package leetcode

func dailyTemperatures2(T []int) []int {
	var res = make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for fd := i + 1; fd < len(T); fd++ {
			if T[fd] > T[i] {
				res[i] = fd - i
				break
			}
		}
	}
	return res
}
