package main

func maxOf(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func minOf(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}
