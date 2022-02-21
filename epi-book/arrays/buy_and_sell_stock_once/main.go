package epi

// Prices : []int{310, 315, 275, 295, (260), 270, (290), 230, 255, 250} Max profit: 30
func maxProfit(p []int) int {
	res := 0
	if len(p) == 1 {
		return res
	}
	lp, rp := 0, 1

	for rp < len(p) {
		if p[lp] < p[rp] {
			res = max(res, p[rp]-p[lp])
			rp++
			continue
		}

		lp = rp
		rp++

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
