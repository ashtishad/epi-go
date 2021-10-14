package main

import "sort"

func threeSums(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			second, third := nums[left], nums[right]
			if first+second+third == 0 {
				res = append(res, []int{first, second, third})
				left++
				for left < right-1 && nums[left] == nums[left-1] {
					left++
				}
				right--
				for right > left && nums[right] == nums[right+1] {
					right--
				}

			} else if first+second+third > 0 {
				right--
			} else {
				left++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}
