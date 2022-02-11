package leetcode

// Input: temperatures = [73,74,(75),71,69,72,(76),73,60,50,40]
// Output: [1,1,4,2,1,1,0,0,0,0,0]
func dailyTemperatures(T []int) []int {
	var n = len(T)
	var res = make([]int, n) // all 0 initialized
	for i := n - 2; i >= 0; i-- {
		var fd = i + 1
		for fd < n {
			// warmer found
			if T[fd] > T[i] {
				res[i] = fd - i
				break
			}
			// no warmer day ahead(case 73 to 40)
			if res[fd] == 0 {
				break
			}
			// T[fd] <= T[i], we can easily use fd result
			// i = 3, 71 has warmer day after 2 days.
			// so, in case of 75(2),as 75 > 71, we can use 71 result
			// and, fast forward this 2 days
			fd = res[fd] + fd
		}
	}
	return res
}
