// Package array /*
/*
	Contains Some array exercises, easy one
*/
package array

type ArrayProblems struct {
}

func NewArrayProblems() *ArrayProblems {
	return &ArrayProblems{}
}

// TwoSum /**
/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
*/
func (p *ArrayProblems) TwoSum(nums []int, target int) []int {
	// requirement => nums (array), target -> int
	// outcome => index of sum of array element that is equal to the target
	// breaking it down => array, target
	// 1st component => constraint validation (length of array must be greater than or equal to 2 and less than or equal to 104 => DONE
	// target must >= -109 && <= 109 same with all element of the array => DONE
	// 2nd component => compare the first item in the array against the target => DONE
	// 3rd component => sum up all items in the array then compare it against the target => DONE
	// 4th component => sum up first item with second, and compare against the target, continue until all items in the array has been used
	// recursion or backtracking here
	// 1 + 2, 1+3, 1+4..... 1+ (n-1)
	// 2 + 3, 2+4 ... 2 + (n-1)
	// n +  (n-1)
	// after each iteration,remove that particular item from the array, also please take note of the index
	// if the condition is still false, the function panic meaning no element can be sum up to be equal to the target

	filter := func(num int) bool {
		return num >= -109 && num <= 109
	}
	validate_num := func(items []int) bool {
		valid_value := make([]int, len(items))
		for i, v := range items {
			if filter(v) {
				valid_value[i] = items[i]
			}
		}

		return len(valid_value) == len(items)
	}
	reduce := func(items []int) int {
		_reduce := func(prev, current int, fn func(a, b int) int) int {
			sum := fn(prev, current)
			return sum
		}
		final_value := 0
		for _, v := range items {
			final_value += _reduce(0, v, func(a, b int) int {
				return a + b
			})
		}

		return final_value
	}

	array_len := len(nums)
	validated := array_len >= 2 && array_len <= 104 && (target >= -109 && target <= 109) && validate_num(nums)
	if !validated {
		panic("constraint validation failed")
	}

	success := (nums[0] == target) || (reduce(nums) == target)
	if success {
		return []int{0}
	}

	answer := make([]int, 0)
	for i := 0; i < array_len; i++ {
		current_index := i
		current_answer := RecursiveSum(nums[i:], 0, i+1, target, make([]int, 0))

		if len(current_answer) > 0 {
			answer = []int{current_index, current_answer[1]}
			break
		}
	}

	return answer
}

func RecursiveSum(items []int, current_index, next_index, expected_value int, index []int) []int {
	if len(items) == 0 {
		return index
	}

	first := items[0]
	next_item := items[next_index]
	sum_item := first + next_item

	if sum_item == expected_value {
		return []int{0, next_index}
	}

	latest_items := shift(items)
	return RecursiveSum(latest_items, current_index, next_index+1, expected_value, index)
}

func shift(items []int) []int {
	if len(items) < 2 {
		return make([]int, 0)
	}

	return items[1:]
}
