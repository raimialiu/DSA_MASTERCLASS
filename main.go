package main

import (
	"fmt"
	"reflection/problems/easy/array"
)

func main() {
	arr := array.NewArrayProblems()

	two_sums := arr.TwoSum([]int{3, 2, 4}, 6)

	fmt.Printf("final answer is %v\n", two_sums)
}

func countdown(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("value is %d\n", i)
	}
}
