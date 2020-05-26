package algorithms

import "math"

///////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Problem 1:                                                                                                //
// Given an array of integers, return indices of the two numbers such that they add up to a specific target. //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
func two_sum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/////////////////////
// Problem 7:	   //
// Reverse Integer //
/////////////////////
func reverse(x int) int {
	var result int64

	for {
		n := x % 10
		x = x / 10

		result = int64(result)*10 + int64(n)

		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}

		if x == 0 {
			break
		}
	}

	return int(result)
}
