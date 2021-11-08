package leetcode

import (
	"fmt"
	"testing"
)

type question121 struct {
	para121
	ans121
}

type para121 struct {
	one []int
}

type ans121 struct {
	one int
}

func Test_Problem121(t *testing.T) {

	qs := []question121{

		{
			para121{[]int{}},
			ans121{0},
		},

		{
			para121{[]int{7, 1, 5, 3, 6, 4}},
			ans121{5},
		},

		{
			para121{[]int{7, 6, 4, 3, 1}},
			ans121{0},
		},

		{
			para121{[]int{1, 3, 2, 8, 4, 9}},
			ans121{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 121------------------------\n")

	for _, q := range qs {
		_, p := q.ans121, q.para121
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit(p.one))
	}
	fmt.Printf("\n\n\n")
}
