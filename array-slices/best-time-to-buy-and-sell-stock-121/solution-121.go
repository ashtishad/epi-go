package leetcode

func maxProfit(prices []int) int {
	res := 0
	if len(prices) == 1 {
		return res
	}
	lp, rp := 0, 1

	for rp < len(prices) {
		if prices[lp] < prices[rp] {
			profit := prices[rp] - prices[lp]
			res = max(res, profit)
			rp++
		} else {
			lp = rp
			rp++
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
