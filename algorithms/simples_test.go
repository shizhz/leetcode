package algorithms

import (
	"fmt"
	"math"
	"reflect"
	"testing"
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
		return fmt.Errorf("Different size")
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
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
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
