package algorithms

import (
	"bytes"
	"math"
	"sort"
	"strings"
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
type charType int

const (
	start charType = iota
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

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Problem 9:																																    //
// Palindrome Number																														    //
// Solution: reverse the number and compare the result with the original one, use int64 to store the reserved result to avoid overflow problem //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	var reversedNumber int64
	for {
		n := x % 10
		x = x / 10

		reversedNumber = reversedNumber*10 + int64(n)

		if x == 0 {
			break
		}
	}

	return reversedNumber == int64(original)
}

// Improvement: just reverse half of the number, and compare the result with the other half
// How do we know we that we've reached the half of the number: consider the following loop:
// - n = x % 10
// - x = x / 10
// - reversedNumber = reversedNumber * 10 + n
//
// when reversedNumber is greater than x for the first time, we may:
// - Just at the right middle point
// - Have done one more step, and crossed the middle point
//
// Then we move one step back, and compare:
// - x == n ?
// - x / 10 == n ? (for odd numbers, we need to remove the middle digit)
func isPalindromeNumberImproved(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 { // A number with ending 0's can not be a palindrome number
		return false
	}

	reversedNumber, left := 0, x
	for {
		n := left % 10
		x = left / 10

		rn := reversedNumber*10 + n

		if rn >= x {
			return reversedNumber == left || reversedNumber == x
		}

		reversedNumber, left = rn, x
	}
}

//////////////////////////////////////////////////////
// Problem 11:									    //
// Container with most water					    //
// Solution: max(abs(i-j) * min(I, J)) for all i, j //
//////////////////////////////////////////////////////
func abs(r int) int {
	if r >= 0 {
		return r
	}

	return -r
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func maxArea(height []int) int {
	var max int

	for i := 0; i < len(height)-1; i++ {
		for j := i; j < len(height); j++ {
			if area := abs(i-j) * min(height[i], height[j]); area > max {
				max = area
			}
		}
	}

	return max
}

//////////////////////
// Problem 12:	    //
// Integer to Roman //
//////////////////////
var romanWeights []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanNumberToSymbol map[int]string = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	var roman bytes.Buffer

	for i := 0; i < len(romanWeights); i++ {
		weight := romanWeights[i]
		for j := 0; j < num/weight; j++ {
			roman.WriteString(romanNumberToSymbol[weight])
		}
		num = num % weight
	}

	return roman.String()
}

//////////////////////
// Problem 13:	    //
// Roman to Integer //
//////////////////////
var romanSymbols []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var romanSymbolToNumber map[string]int = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	var result int

	for _, symbol := range romanSymbols {
		for strings.HasPrefix(s, symbol) {
			result += romanSymbolToNumber[symbol]
			s = s[len(symbol):]
		}
	}

	return result
}

///////////////////////////
// Problem 14:			 //
// Longest Common Prefix //
///////////////////////////
func lcp(s1, s2 string) string {
	var prefix bytes.Buffer
	for i := 0; ; i++ {
		if i >= len(s1) || i >= len(s2) {
			break
		}

		if s1[i] == s2[i] {
			prefix.WriteByte(s1[i])
			continue
		}

		break
	}

	return prefix.String()
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string = strs[0]

	for i := 1; i < len(strs); i++ {
		p := lcp(prefix, strs[i])
		if len(p) == 0 {
			return ""
		}

		prefix = p
	}

	return prefix
}

/////////////////////////////////////////////////////////////////////////////
// Problem 15:															   //
// 3Sum: Find all unique triplets in the array which gives the sum of zero //
/////////////////////////////////////////////////////////////////////////////
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	// fmt.Printf("Array : %v\n", nums)

	var result [][]int = [][]int{}

	if len(nums) < 3 {
		return result
	}

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}

	end := len(nums) - 2

	for i := 0; i < end; {
		in := nums[i]
		if in > 0 {
			break
		}
		j := len(nums) - 1

		for {
			jn := nums[j]
			if j <= i+1 || jn < 0 {
				break
			}

			theSecondOne := -(in + jn)

			index := indexOf(nums, i+1, j, theSecondOne)
			// fmt.Printf("i: %d[%d], j: %d[%d], try to find: %d. Found index: %d", i, in, j, jn, theSecondOne, index)
			if index < j && nums[index] == theSecondOne {
				// fmt.Println(" Fount It")
				result = append(result, []int{in, theSecondOne, jn})
			}

			/* Moving by simply decreasing */
			j--

			for nums[j] == nums[j+1] && j > i+1 {
				j--
			}

			// Move right part to left, Use binary search to skip the continous ones
			// if nums[j-1] == jn {
			// 	j = indexOf(nums, i+1, j, jn) - 1
			// 	// fmt.Printf("Next J:%d by binary moving\n", j)
			// } else {
			// 	j--
			// }
		}

		/* Moving by simply decreasing */
		i++

		for nums[i] == nums[i-1] && i < end {
			i++
		}

		// if nums[i+1] == in {
		// 	i = indexOf(nums, i, len(nums), in+1) // Move left part to right, Use binary search to skip the continous ones
		// } else {
		// 	i++
		// }
	}
	return result
}

////////////////////////////////////////////////////////////////////////////////////////////
// Problem 16: 																		    //
// 3Sum closest                                                                           //
////////////////////////////////////////////////////////////////////////////////////////////
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	if len(nums) < 3 {
		return 0
	}

	var result int = nums[0] + nums[1] + nums[2]

	tryNumbersAt := func(i, j, k int) {
		sum := nums[i] + nums[j] + nums[k]
		if abs(sum-target) < abs(result-target) {
			result = sum
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := len(nums) - 1; j > i+1; j-- {
			theSecondOne := target - (nums[i] + nums[j]) // The pefect candidate number

			index := indexOf(nums, i+1, j, theSecondOne)
			// fmt.Printf("i: %d[%d], j: %d[%d], try to find: %d. Found index: %d", i, in, j, jn, theSecondOne, index)
			if index == j {
				tryNumbersAt(i, j, index-1)
			} else if index == i+1 {
				tryNumbersAt(i, j, index)
			} else {
				tryNumbersAt(i, j, index)
				tryNumbersAt(i, j, index-1)
			}

			if result == target {
				return result
			}
		}
	}
	return result
}
