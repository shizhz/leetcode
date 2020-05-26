package algorithms

import (
	"math"
)

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

/////////////////////////////
// Problem 8:			   //
// String to Integer(atoi) //
/////////////////////////////
type iotaState int

const (
	start iotaState = iota
	whitespace
	sign
	number
)

func myAtoi(s string) int {
	ws, minus, plus := rune(32), rune(45), rune(43)
	var result int64
	var resultSign int64 = 1
	state := start

	for _, c := range s {
		switch {
		case c == ws:
			switch {
			case state == whitespace:
				// nop, ignore leading whitespaces
			case state == start:
				state = whitespace
			default:
				goto done
			}
		case c == minus || c == plus:
			switch {
			case state == whitespace, state == start:
				state = sign
				if c == minus {
					resultSign = -1
				}
			default:
				goto done
			}
		case c >= 48 && c <= 57:
			switch {
			case state == whitespace, state == start, state == sign, state == number:
				state = number
				result = result*int64(10) + resultSign*(int64(c)-48)

				if result > math.MaxInt32 {
					return math.MaxInt32
				}

				if result < math.MinInt32 {
					return math.MinInt32
				}
			default:
				goto done
			}
		default:
			goto done
		}
	}

done:
	return int(result)
}
