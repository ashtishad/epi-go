package leetcode

import (
	"fmt"
	"testing"
)

type question136 struct {
	para136
	ans136
}

type para136 struct {
	s []int
}

type ans136 struct {
	one int
}

func Test_Problem136(t *testing.T) {

	qs := []question136{

		{
			para136{[]int{2, 2, 1}},
			ans136{1},
		},

		{
			para136{[]int{4, 1, 2, 1, 2}},
			ans136{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 136------------------------\n")

	for _, q := range qs {
		_, p := q.ans136, q.para136
		fmt.Printf("【input】:%v       【output】:%v\n", p, singleNumber(p.s))
	}
	fmt.Printf("\n\n\n")
}
