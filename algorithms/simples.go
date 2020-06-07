package algorithms

import (
	"bytes"
	"fmt"
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
func threeSumToTarget(nums []int, target int) [][]int {
	sort.Ints(nums)

	// fmt.Printf("Array : %v\n", nums)

	var result [][]int = [][]int{}

	if len(nums) < 3 {
		return result
	}

	end := len(nums) - 2

	for i := 0; i < end; {
		in := nums[i]
		j := len(nums) - 1

		for {
			jn := nums[j]
			if j <= i+1 {
				break
			}

			theSecondOne := target - (in + jn)

			index := indexOf(nums, i+1, j, theSecondOne)
			// fmt.Printf("%v, i: %d[%d], j: %d[%d], try to find: %d. Found index: %d\n", nums, i, in, j, jn, theSecondOne, index)
			if index < j && nums[index] == theSecondOne {
				result = append(result, []int{in, theSecondOne, jn})
			}

			/* Moving by simply decreasing */
			j--

			for nums[j] == nums[j+1] && j > i+1 {
				j--
			}
		}

		/* Moving by simply decreasing */
		i++

		for nums[i] == nums[i-1] && i < end {
			i++
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	return threeSumToTarget(nums, 0)
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

///////////////////////////////////////////
// Problem 17:							 //
// Letter Combinations of a Phone Number //
///////////////////////////////////////////
var digitToLetter map[rune]string = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func combLetter(combs []string, letter byte) []string {
	if len(combs) == 0 {
		return []string{string(letter)}
	}

	var result []string = []string{}

	for _, com := range combs {
		result = append(result, com+string(letter))
	}

	return result

}

func comb(combs []string, letters string) []string {
	if len(letters) == 0 {
		return []string{}
	}

	return append(combLetter(combs, letters[0]), comb(combs, letters[1:])...)
}

func letterCombinations(digits string) []string {
	var result []string = []string{}

	for _, digit := range digits {
		result = comb(result, digitToLetter[digit])
	}

	return result
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Problem 18:																																	 //
// 4Sum: Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target?	 //
// 	  Find all unique quadruplets in the array which gives the sum of target.																	 //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int = [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		// skip same numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		subCombinations := threeSumToTarget(nums[i+1:], target-nums[i])

		// fmt.Printf("Nums: %v, Got Coms: %v. nums: %v, target: %d\n", nums, subCombinations, nums[i+1:], target-nums[i])

		for _, c := range subCombinations {
			com := append(c, nums[i])
			sort.Ints(com)
			result = append(result, com)
		}
	}

	return result
}

//////////////////////////////////////////////////
// Problem 19: Remove Nth Node From End of List //
//////////////////////////////////////////////////

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) String() string {
	ints := []int{}

	for this != nil {
		ints = append(ints, this.Val)
		this = this.Next
	}

	return fmt.Sprint(ints)
}

func listNodeSize(this *ListNode) int {
	size := 0

	for this != nil {
		size++
		this = this.Next
	}

	return size
}

// Stratforward method:
// 1. Loop one time to count how many nodes the list has
// 2. Remove the target node
func removeNthFromEndTwoLoops(head *ListNode, n int) *ListNode {
	size := listNodeSize(head)

	h := &ListNode{
		Next: head,
	}

	node := h

	for i := 0; i < size-n; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next

	return h.Next
}

// One loop algorithm: use two pointer P1 and P2, and make indexOf(P2) - indexOf(P1) = n + 1, thus when P2 reaches the end(nil), P1 will be
// the previous index of the target node. Then make P1.Next = P1.Next.Next and return head
func removeNthFromEndOneLoop(head *ListNode, n int) *ListNode {
	h := &ListNode{
		Next: head,
	}

	p1, p2 := h, h
	step := 0

	for p2 != nil {
		p2 = p2.Next
		if step == n+1 {
			p1 = p1.Next
		} else {
			step++
		}
	}

	p1.Next = p1.Next.Next

	return h.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// return removeNthFromEndTwoLoops(head, n)
	return removeNthFromEndOneLoop(head, n)
}

///////////////////////////////////
// Problem 20: Valid Parentheses //
///////////////////////////////////
func isValid(s string) bool {
	stack := []rune{}

	match := func(c rune) bool {
		// Improment thought: the ascii code for each char is:
		// - ( : 40
		// - ) : 41
		// - [ : 91
		// - ] : 93
		// - { : 123
		// - } : 125
		// Consider the input only contains those chars, so if `c` - stack[len(stack) - 1] == 1 or 2, it must be a valid pair
		if len(stack) == 0 {
			return false
		}
		t := stack[len(stack)-1]
		switch c {
		case ']':
			return t == '['
		case '}':
			return t == '{'
		case ')':
			return t == '('
		default:
			return false
		}
	}

	isOpen := func(c rune) bool {
		return c == '[' || c == '{' || c == '('
	}

	// isClose := func(c rune) bool {
	// 	return c == ']' || c == '}' || c == ')'
	// }

	for _, c := range s {
		if isOpen(c) {
			stack = append(stack, c)
		} else if match(c) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// Problem 21: Merge Two Sorted Lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				node.Next = l1
				l1 = l1.Next
			} else {
				node.Next = l2
				l2 = l2.Next
			}
		} else if l1 == nil {
			node.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			node.Next = l1
			l1 = l1.Next
		}

		node = node.Next
	}

	return head.Next
}

// Problem 22: Generate parentheses
func doGenerateParenthesis(rcv chan string, s string, left, right int) {
	if left == 0 && right == 0 {
		rcv <- s
		return
	}

	if left == 0 {
		doGenerateParenthesis(rcv, s+")", left, right-1)
		return
	}

	if left < right {
		doGenerateParenthesis(rcv, s+"(", left-1, right)
		doGenerateParenthesis(rcv, s+")", left, right-1)
	} else {
		doGenerateParenthesis(rcv, s+"(", left-1, right)
	}
}

func generateParenthesis(n int) []string {
	ch := make(chan string)

	result := []string{}

	go func() {
		defer close(ch)
		doGenerateParenthesis(ch, "", n, n)
	}()

	for s := range ch {
		result = append(result, s)
	}

	return result
}

// Problem 23 Merge k Sorted Lists
func mergeKLists(lists []*ListNode) *ListNode {
	heads := make([]*ListNode, len(lists))
	for i := 0; i < len(lists); i++ {
		heads[i] = lists[i]
	}

	var minIndex int = -1
	var minValue int

	result := &ListNode{}
	node := result

	for {
		for index, head := range heads {
			if head != nil {
				if minIndex == -1 || head.Val < minValue {
					minIndex, minValue = index, head.Val
				}
			}
		}
		if minIndex == -1 {
			break
		}

		node.Next = &ListNode{
			Val: minValue,
		}
		node = node.Next
		heads[minIndex] = heads[minIndex].Next
		minIndex = -1
	}

	return result.Next
}

// Problem 24: Swap Nodes in Pairs
func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{
		Next: head,
	}

	prev := result

	for {
		if head == nil || head.Next == nil {
			break
		}

		temp := head.Next
		head.Next = temp.Next
		temp.Next = head
		prev.Next = temp

		prev = head
		head = head.Next
	}

	return result.Next
}

// Problem 25: Reverse Nodes in k-Group
// Rotate the list of range [i, j] for one element, e.g, if input are:
// - list : [1,2,3,4,5]
// - i: 2
// - j: 4
// then the output becomes: [1, 2, 5, 3, 4]
func rotateList(head *ListNode, i, j int) *ListNode {
	if i == j {
		return head
	}
	if i > j {
		i, j = j, i
	}

	result := &ListNode{
		Next: head,
	}

	previ := result
	prevj := result

	for k := 0; k < j; k++ {
		if previ == nil || prevj == nil {
			return head
		}
		if k < i {
			previ = previ.Next
		}
		prevj = prevj.Next
	}

	nodei := previ.Next
	nodej := prevj.Next

	prevj.Next = nodej.Next
	previ.Next = nodej

	nodej.Next = nodei

	return result.Next
}

func reverseTopNOfList(head *ListNode, n int) (*ListNode, bool) {
	if n <= 1 {
		return head, false
	}

	tail := head

	// Make sure there're enough elements left, otherwise don't do anything
	for i := 0; i < n; i++ {
		if tail == nil {
			return head, false
		}
		tail = tail.Next
	}

	prevNode := head
	node := head.Next
	head.Next = tail

	for i := 1; i < n; i++ {
		temp := node.Next
		node.Next = prevNode
		prevNode = node
		node = temp
	}

	return prevNode, true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	resultHead, reversed := reverseTopNOfList(head, k)

	subHead := resultHead

	for reversed {
		for i := 0; i < k-1; i++ {
			subHead = subHead.Next
		}

		if reversedHead, yes := reverseTopNOfList(subHead.Next, k); yes {
			subHead.Next = reversedHead
			subHead = subHead.Next
		} else {
			reversed = false
		}
	}

	return resultHead
}

// 26. Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	var index int

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[index] {
			continue
		}

		nums[index+1] = nums[i]
		index++
	}

	return index + 1
}

// 27. Remove Element
func removeElement(nums []int, val int) int {
	indices := []int{}
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			indices = append(indices, i)
		} else {
			size++
			if len(indices) > 0 {
				nums[indices[0]] = nums[i]
				indices = append(indices[1:], i)
			}
		}
	}

	return size
}

// 28. Implement strStr()
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle) // hahahahahahahaha
}

// 29. Divide Two Integers
// The key is using substraction, after handling the sign seperately, for the case of two positive numbers,
// the number of times we can substract divisor from dividend, is the result of quotient
// 该题的核心思想是使用减法，在对符号进行过处理，只考虑两个正数相除的前提下，能够从dividend中减去多少个divisor，那么商就是多少
//
// And the optimization is trying to substract divisor from dividend as fast as possiable
// 优化的方式便是，想办法尽可能快的从dividend中减去足够的divisor

// Implementation 1: Substract only one divisor from dividend each loop, this is the most straightforward but also the worst method
// 实现方案一：每次从dividend中减去一个divisor，直到不能减为止，这是最直观也是最差的方案
func divideBySubstraction(absDividend uint64, absDivisor uint64) int {

	var result int

	for {
		if absDividend-absDivisor >= absDivisor {
			absDividend -= absDivisor
			result++
			continue
		}
		break
	}

	return result + 1
}

// Implementation 2: With the help of bit shift operation, we can test whether (dividend>>1 > divisor),
// if so, we know this is true: divident > divisor * 2, so we can substract two divisors from dividend for this loop
// if not, we know we can substract at most one more divisor from dividend
// 实现方案二：借助于移位操作，我们可以测试dividend>>1之后是否还大于divisor，dividend>>1相当于dividend / 2， 这样的话说明dividend > divisor * 2是成立的，
// 根据该不等式，我们无法知道1/2的dividend最多可以减去多少个divisor，但是我们可以知道至少减去2个divisor是可以的；如果dividend>>1已经小于divisor，那么可知
// 最多只能减去一个divisor了
//
// 事实证明：这个实现比纯粹使用减法还差劲！！
func divideByBitShiftAndSubstraction(absDividend uint64, absDivisor uint64) int {
	var result int

	for {
		if (absDividend >> 1) > absDivisor {
			result += 2
			absDividend = absDividend - absDivisor<<1
			continue
		}
		break
	}
	return result + divideBySubstraction(absDividend, absDivisor)
}

// Implementation 3: Let's try reversed thinking and see theoretically at most how many advisors we can substrct from dividend for each loop:
// Say the final quotient has binary reprentation of: 1011011011, which is 2^9 + 2^7 + 2^6 + 2^4 + 2^3 + 2^1 + 2^0,
// then how could we find a way to substract divisors from dividend by the order of (2^9, 2^7, 2^6, 2^4, 2^3, 2^1, 2^0)?
// In implementation 2 we try bit-shift with one bit each time, but if we try bit-shift with this order (31, 30, 29....0), for all the numbers n,
// which satisfies dividend >> n >= divisor, we know we can safely substract 2^n divisors from dividend and continue the loop
// 实现方案三：回到最初的思路上来：如何每次减去最多的advisor？在方法二中我们每次移一个位，所以每次最多只能减去2个divisor，考虑到实际上被除数与除数是32位整数，那么可以由高到底地
// 对被除数进行移位（31,30,29....0），对于每一个dividend >> n > divisor, 我们都可以安全地减掉2^n个divisor
func divideByBitShiftOnly(absDividend, absDivisor uint64) int {
	var result int

	for i := 31; i >= 0; i-- {
		if absDividend>>i >= absDivisor {
			absDividend -= absDivisor << i
			result += (1 << i)
		}
	}

	return result

}

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	// Overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := dividend^divisor < 0

	absDividend := uint64(math.Abs(float64(dividend)))
	absDivisor := uint64(math.Abs(float64(divisor)))

	if absDividend < absDivisor {
		return 0
	}

	// q := divideBySubstraction(absDividend, absDivisor)
	// q := divideByBitShiftAndSubstraction(absDividend, absDivisor)
	q := divideByBitShiftOnly(absDividend, absDivisor)
	if sign {
		return -q
	}

	return q
}
