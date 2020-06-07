package algorithms

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_two_sum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 01",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "test 02",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 18,
			},
			want: []int{1, 2},
		},
		{
			name: "test 03",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 17,
			},
			want: []int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := two_sum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("two_sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "test 01",
			args: 123,
			want: 321,
		},
		{
			name: "test 02",
			args: -123,
			want: -321,
		},
		{
			name: "test 03",
			args: 120,
			want: 21,
		},
		{
			name: "test 04",
			args: math.MaxInt32,
			want: 0,
		},
		{
			name: "test 05",
			args: math.MinInt32,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "test 01",
			args: "42",
			want: 42,
		},
		{
			name: "test 02",
			args: "  -42",
			want: -42,
		},
		{
			name: "test 03",
			args: "4193 with words ",
			want: 4193,
		},
		{
			name: "test 04",
			args: "words and 987",
			want: 0,
		},
		{
			name: "test 05",
			args: "-91283472332",
			want: -2147483648,
		},
		{
			name: "test 06",
			args: "91283472332",
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args); got != tt.want {
				t.Errorf("myatoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeNumber(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "test 01",
			args: 121,
			want: true,
		},
		{
			name: "test 02",
			args: 1211,
			want: false,
		},
		{
			name: "test 03",
			args: -121,
			want: false,
		},
		{
			name: "test 04",
			args: 0,
			want: true,
		},
		{
			name: "test 05",
			args: 10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNumber(tt.args); got != tt.want {
				t.Errorf("isPalindromeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeNumberImproved(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "test 01",
			args: 121,
			want: true,
		},
		{
			name: "test 02",
			args: 1211,
			want: false,
		},
		{
			name: "test 03",
			args: -121,
			want: false,
		},
		{
			name: "test 04",
			args: 0,
			want: true,
		},
		{
			name: "test 05",
			args: 10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNumberImproved(tt.args); got != tt.want {
				t.Errorf("isPalindromeNumberImproved() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.r); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test 01",
			args: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "test 01",
			args: 0,
			want: "",
		},
		{
			name: "test 02",
			args: 1,
			want: "I",
		},
		{
			name: "test 03",
			args: 3,
			want: "III",
		},
		{
			name: "test 04",
			args: 4,
			want: "IV",
		},
		{
			name: "test 05",
			args: 4,
			want: "IV",
		},
		{
			name: "test 06",
			args: 9,
			want: "IX",
		},
		{
			name: "test 07",
			args: 58,
			want: "LVIII",
		},
		{
			name: "test 08",
			args: 1994,
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcp(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 01",
			args: args{
				s1: "a",
				s2: "b",
			},
			want: "",
		},
		{
			name: "test 02",
			args: args{
				s1: "",
				s2: "",
			},
			want: "",
		},
		{
			name: "test 03",
			args: args{
				s1: "abc",
				s2: "a",
			},
			want: "a",
		},
		{
			name: "test 04",
			args: args{
				s1: "abc",
				s2: "abd",
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcp(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("lcp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 01",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "test 02",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "test 01",
			args: "III",
			want: 3,
		},
		{
			name: "test 02",
			args: "IV",
			want: 4,
		},
		{
			name: "test 03",
			args: "IX",
			want: 9,
		},
		{
			name: "test 04",
			args: "LVIII",
			want: 58,
		},
		{
			name: "test 05",
			args: "MCMXCIV",
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_threeSumToTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 01",
			args: args{
				nums:   []int{0, 1, 2},
				target: 3,
			},
			want: [][]int{{0, 1, 2}},
		},
		{
			name: "test 02",
			args: args{
				nums:   []int{-2, 1, 1, 3, 5, 5, 5},
				target: 9,
			},
			want: [][]int{{1, 3, 5}},
		},
		{
			name: "test 03",
			args: args{
				nums:   []int{-4, -3, -2, 1, 3, 3, 5},
				target: -6,
			},
			want: [][]int{{-4, -3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumToTarget(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "test 01",
			args: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "test 02",
			args: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "test 03",
			args: []int{-2, 0, 0, 2, 2},
			want: [][]int{
				{-2, 0, 2},
			},
		},
		{
			name: "test 04",
			args: []int{-3, -3, 0, -5},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 01",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "test 02",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: 0,
			},
			want: 3,
		},
		{
			name: "test 03",
			args: args{
				nums:   []int{1, 1, 1, 0},
				target: 100,
			},
			want: 3,
		},
		{
			name: "test 04",
			args: args{
				nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
				target: 82,
			},
			want: 82,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combLetter(t *testing.T) {
	type args struct {
		combs  []string
		letter byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test 01",
			args: args{
				combs:  []string{},
				letter: 'a',
			},
			want: []string{"a"},
		},
		{
			name: "test 02",
			args: args{
				combs:  []string{"a", "b", "c"},
				letter: 'a',
			},
			want: []string{"aa", "ba", "ca"},
		},
		{
			name: "test 03",
			args: args{
				combs:  []string{"aa", "ba"},
				letter: 'a',
			},
			want: []string{"aaa", "baa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combLetter(tt.args.combs, tt.args.letter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_comb(t *testing.T) {
	type args struct {
		combs   []string
		letters string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test 01",
			args: args{
				combs:   []string{},
				letters: "abc",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "test 02",
			args: args{
				combs:   []string{"a"},
				letters: "abc",
			},
			want: []string{"aa", "ab", "ac"},
		},
		{
			name: "test 03",
			args: args{
				combs:   []string{"a", "b", "cc"},
				letters: "abc",
			},
			want: []string{"aa", "ab", "ac", "ba", "bb", "bc", "cca", "ccb", "ccc"},
		},
		{
			name: "test 04",
			args: args{
				combs:   []string{"a", "b", "cc"},
				letters: "",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := comb(tt.args.combs, tt.args.letters)

			if err := stringSliceEquals(got, tt.want); err != nil {
				t.Errorf("comb() = %v, want %v. Reason: %s", got, tt.want, err.Error())
			}
		})
	}
}

func stringSliceEquals(s1, s2 []string) error {
	s1Map := map[string]bool{}
	for _, val := range s1 {
		s1Map[val] = true
	}

	if len(s2) != len(s1Map) {
		return fmt.Errorf("Different size. s1: %d, s2: %d", len(s1Map), len(s2))
	}

	for _, ele := range s2 {
		if !s1Map[ele] {
			return fmt.Errorf("`%s` in the second slice but not in the first one", ele)
		}
	}

	return nil
}

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "test 01",
			args: "23",
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args)
			if err := stringSliceEquals(got, tt.want); err != nil {
				t.Errorf("letterCombinations() = %v, want %v. Reason: %s", got, tt.want, err.Error())
			}
		})
	}
}

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 01",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{
				{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
			},
		},
		{
			name: "test 02",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			name: "test 03",
			args: args{
				nums:   []int{2, 1, 0, -1},
				target: 2,
			},
			want: [][]int{
				{-1, 0, 1, 2},
			},
		},
		{
			name: "test 04",
			args: args{
				nums:   []int{5, 5, 3, 5, 1, -5, 1, -2},
				target: 4,
			},
			want: [][]int{
				{-5, 1, 3, 5},
			},
		},
		{
			name: "test 04",
			args: args{
				nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
				target: -11,
			},
			want: [][]int{
				{-5, -4, -3, 1},
			},
		},
		{
			name: "test 05",
			args: args{
				nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
				target: -11,
			},
			want: [][]int{
				{-5, -4, -3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeNodeList(nums []int) *ListNode {
	head := &ListNode{}
	result := head

	for i := 0; i < len(nums); i++ {
		head.Next = &ListNode{
			Val: nums[i],
		}
		head = head.Next
	}

	return result.Next
}

func Test_removeNthFromEnd(t *testing.T) {
	// list := makeNodeList([]int{1, 2, 3, 4, 5})

	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: makeNodeList([]int{1, 2, 3, 5}),
		},
		{
			name: "test 02",
			args: args{
				head: makeNodeList([]int{1}),
				n:    1,
			},
			want: makeNodeList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_len(t *testing.T) {
	assert.Equal(t, 5, listNodeSize(makeNodeList([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, 0, listNodeSize(makeNodeList([]int{})))
}

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "test 01",
			args: "[]",
			want: true,
		},
		{
			name: "test 02",
			args: "[]{}",
			want: true,
		},
		{
			name: "test 03",
			args: "[{]}",
			want: false,
		},
		{
			name: "test 04",
			args: "[{}]",
			want: true,
		},
		{
			name: "test 05",
			args: "()[]{}[{}]",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				l1: makeNodeList([]int{1, 2, 4}),
				l2: makeNodeList([]int{1, 3, 4}),
			},
			want: makeNodeList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "test 02",
			args: args{
				l2: makeNodeList([]int{1, 3, 4}),
			},
			want: makeNodeList([]int{1, 3, 4}),
		},
		{
			name: "test 03",
			args: args{
				l1: makeNodeList([]int{1, 3, 4}),
			},
			want: makeNodeList([]int{1, 3, 4}),
		},
		{
			name: "test 04",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		{
			name: "test 01",
			args: 3,
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			name: "test 02",
			args: 4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args)
			sort.Strings(got)
			sort.Strings(tt.want)
			fmt.Printf("Got: %s, Size: %d\nExp: %s, Size: %d\n", got, len(got), tt.want, len(tt.want))
			if err := stringSliceEquals(got, tt.want); err != nil {
				t.Errorf("generateParenthesis() = %v, want %v. Reason: %s", got, tt.want, err)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				lists: []*ListNode{
					makeNodeList([]int{1, 4, 5}),
					makeNodeList([]int{1, 3, 4}),
					makeNodeList([]int{2, 6}),
				},
			},
			want: makeNodeList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "test 02",
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			name: "test 03",
			args: args{
				lists: []*ListNode{
					makeNodeList([]int{2, 6}),
				},
			},
			want: makeNodeList([]int{2, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4}),
			},
			want: makeNodeList([]int{2, 1, 4, 3}),
		},
		{
			name: "test 02",
			args: args{
				head: makeNodeList([]int{1, 2, 3}),
			},
			want: makeNodeList([]int{2, 1, 3}),
		},
		{
			name: "test 02",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateList(t *testing.T) {
	type args struct {
		head *ListNode
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				i:    2,
				j:    4,
			},
			want: makeNodeList([]int{1, 2, 5, 3, 4}),
		},
		{
			name: "test 02",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				i:    1,
				j:    4,
			},
			want: makeNodeList([]int{1, 5, 2, 3, 4}),
		},
		{
			name: "test 03",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				i:    1,
				j:    7,
			},
			want: makeNodeList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "test 04",
			args: args{
				head: nil,
				i:    1,
				j:    7,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateList(tt.args.head, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseTopNOfList(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				head: nil,
				n:    1,
			},
			want: nil,
		},
		{
			name: "test 02",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    1,
			},
			want: makeNodeList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "test 03",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: makeNodeList([]int{2, 1, 3, 4, 5}),
		},
		{
			name: "test 04",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    3,
			},
			want: makeNodeList([]int{3, 2, 1, 4, 5}),
		},
		{
			name: "test 05",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    4,
			},
			want: makeNodeList([]int{4, 3, 2, 1, 5}),
		},
		{
			name: "test 06",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    5,
			},
			want: makeNodeList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "test 07",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				n:    6,
			},
			want: makeNodeList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := reverseTopNOfList(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseTopNOfList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test 01",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: makeNodeList([]int{2, 1, 4, 3, 5}),
		},
		{
			name: "test 02",
			args: args{
				head: makeNodeList([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: makeNodeList([]int{3, 2, 1, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test 01",
			args: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "test 02",
			args: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 01",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "test 02",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "test 03",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 01",
			args: args{
				dividend: 100,
				divisor:  3,
			},
			want: 33,
		},
		{
			name: "test 02",
			args: args{
				dividend: -99,
				divisor:  3,
			},
			want: -33,
		},
		{
			name: "test 03",
			args: args{
				dividend: -99,
				divisor:  0,
			},
			want: 0,
		},
		{
			name: "test 04",
			args: args{
				dividend: math.MinInt32,
				divisor:  -1,
			},
			want: math.MaxInt32,
		},
		{
			name: "test 05",
			args: args{
				dividend: 100,
				divisor:  2,
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDivideBySubstraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divideBySubstraction(10000, 3)
	}
}

func BenchmarkDivideByBitShiftAndSubstraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divideByBitShiftAndSubstraction(10000, 3)
	}
}

func BenchmarkDivideByBitShiftOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divideByBitShiftOnly(10000, 3)
	}
}
