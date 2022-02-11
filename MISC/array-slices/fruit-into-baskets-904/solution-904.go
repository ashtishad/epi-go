// length of longest subarray that contains two distinct integers(fruits)
// Dynamic sized window
// https://leetcode.com/problems/fruit-into-baskets/discuss/170745/Problem%3A-Longest-Subarray-With-2-Elements
package leetcode

// [3,3,3,1,2,1,1,2,3,3,4]
func totalFruit(trees []int) int {
	var ctr = make(map[int]bool) // k: fruit type, v: basket exist or not
	var res, start = 0, 0
	for end, v := range trees {
		var ws = end - start + 1 // window size
		// basket available and not in map
		if len(ctr) < 2 && !ctr[v] {
			ctr[v] = true
			res = max(res, ws)
			continue
		}

		// ok to collect fruit into existing basket
		if ctr[v] {
			res = max(res, ws)
			continue
		}

		// starting new window,considering prev fruit and go ahead with current fruit
		// happned at idx 3(1) and idx 4(2)
		ctr = make(map[int]bool) // reset baskets
		ctr[trees[end-1]] = true
		ctr[v] = true
		start = end - 1

		// while resetting window checking if it has contageous duplicates.
		for trees[start] == trees[start-1] {
			start--
		}
		// resetted window, so calculated again
		ws = end - start + 1
		res = max(res, ws)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
