package epi

// Prices : []int{310, 315, 275, 295, (260), 270, (290), 230, 255, 250} Max profit: 30
// Two pointers, slow and fast.
func maxProfit(p []int) int {
	res := 0
	if len(p) == 1 {
		return res
	}
	sp, fp := 0, 1

	for fp < len(p) {
		if p[sp] < p[fp] {
			res = max(res, p[fp]-p[sp])
			fp++
			continue
		}

		sp = fp
		fp++

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
