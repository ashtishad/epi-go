package for_range_loop

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}

	for i := range nums {
		// don't use 2nd variable "num" here
		// it increases complexity
		// use _ instead, it's faster and best practice

		fmt.Println(nums[i]) // access with nums[i]

	}
}
