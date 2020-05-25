package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_bruteforce_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "empty string",
			args: "",
			want: "",
		},
		{
			name: "single char",
			args: "a",
			want: "a",
		},
		{
			name: "no palindrome",
			args: "abcde",
			want: "a",
		},
		{
			name: "test - 01",
			args: "babad",
			want: "bab",
		},
		{
			name: "test - 02",
			args: "cbbd",
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &bruteforce{}
			if got := this.longestPalindrome(tt.args); got != tt.want {
				t.Errorf("bruteforce.longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_bruteforce_isPalindrome(t *testing.T) {
	bf := &bruteforce{}
	tests := []struct {
		name string
		this *bruteforce
		args string
		want bool
	}{
		{
			name: "empty string",
			this: bf,
			args: "",
			want: true,
		},
		{
			name: "even len string true",
			this: bf,
			args: "abba",
			want: true,
		},
		{
			name: "even len string false",
			this: bf,
			args: "abbaaa",
			want: false,
		},
		{
			name: "odd len string true",
			this: bf,
			args: "aba",
			want: true,
		},
		{
			name: "odd len string false",
			this: bf,
			args: "ababb",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("bruteforce.isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bruteforce_substrings(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			name: "empty",
			args: "",
		},
		{
			name: "one char",
			args: "a",
		},
		{
			name: "multiple chars",
			args: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &bruteforce{}
			fmt.Printf("Test substirngs of : %s\n", tt.args)
			for sub := range this.substrings(tt.args) {
				fmt.Println(sub)
			}
			fmt.Printf("Test substirngs of : %s DONE\n", tt.args)
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandCenter_extractPalindrome(t *testing.T) {
	type args struct {
		j int
		k int
	}
	tests := []struct {
		name string
		s    string
		args args
		want string
	}{
		{
			name: "test - 02",
			args: args{
				j: 0,
				k: 0,
			},
			s:    "abcde",
			want: "a",
		},
		{
			name: "test - 03",
			args: args{
				j: 1,
				k: 1,
			},
			s:    "abcde",
			want: "b",
		},
		{
			name: "test - 04",
			args: args{
				j: 1,
				k: 1,
			},
			s:    "abade",
			want: "aba",
		},
		{
			name: "test - 05",
			args: args{
				j: 1,
				k: 1,
			},
			s:    "babade",
			want: "bab",
		},
		{
			name: "test - 06",
			args: args{
				j: 2,
				k: 2,
			},
			s:    "eabade",
			want: "aba",
		},
		{
			name: "test - 07",
			args: args{
				j: 2,
				k: 3,
			},
			s:    "eabade",
			want: "",
		},
		{
			name: "test - 08",
			args: args{
				j: 1,
				k: 2,
			},
			s:    "eaaa",
			want: "aa",
		},
		{
			name: "test - 09",
			args: args{
				j: 2,
				k: 3,
			},
			s:    "eabbac",
			want: "abba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &expandCenter{}
			if got := this.extractPalindrome(tt.s, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("expandCenter.extractPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandCenter_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "empty string",
			args: "",
			want: "",
		},
		{
			name: "single char",
			args: "a",
			want: "a",
		},
		{
			name: "no palindrome",
			args: "abcde",
			want: "a",
		},
		{
			name: "test - 01",
			args: "babad",
			want: "bab",
		},
		{
			name: "test - 02",
			args: "cbbd",
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &expandCenter{}
			if got := this.longestPalindrome(tt.args); got != tt.want {
				t.Errorf("expandCenter.longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "empty string",
			args: "",
			want: "",
		},
		{
			name: "single char",
			args: "a",
			want: "a",
		},
		{
			name: "no palindrome",
			args: "abcde",
			want: "a",
		},
		{
			name: "test - 01",
			args: "babad",
			want: "aba", // bab is also a correct answer, but dp will return aba
		},
		{
			name: "test - 02",
			args: "cbbd",
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &dp{
				memo: map[int]map[int]bool{},
			}
			if got := this.longestPalindrome(tt.args); got != tt.want {
				t.Errorf("dp.longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp_longestPalindromeBetween(t *testing.T) {
	type args struct {
		s string
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 01",
			args: args{
				s: "a",
				i: 0,
				j: 0,
			},
			want: "a",
		},
		{
			name: "test 02",
			args: args{
				s: "aabaa",
				i: 0,
				j: 1,
			},
			want: "aa",
		},
		{
			name: "test 03",
			args: args{
				s: "aabaa",
				i: 0,
				j: 4,
			},
			want: "aabaa",
		},
		{
			name: "test 03",
			args: args{
				s: "aabca",
				i: 0,
				j: 2,
			},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &dp{
				memo: map[int]map[int]bool{},
			}
			if got := this.longestPalindromeBetween(tt.args.s, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("dp.longestPalindromeBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp_memorize(t *testing.T) {
	this := &dp{
		memo: map[int]map[int]bool{
			0: {
				0: true,
				1: false,
				2: true,
			},
			1: {
				9: true,
				6: false,
				2: false,
			},
		},
	}
	want := map[int]map[int]bool{
		0: {
			0: true,
			1: false,
			2: true,
		},
		1: {
			9: true,
			6: false,
			2: false,
		},
	}
	type args struct {
		i            int
		j            int
		isPalindrome bool
	}
	tests := []struct {
		name string
		args []args
	}{
		{
			name: "test - 01",
			args: []args{
				{
					i:            0,
					j:            0,
					isPalindrome: true,
				},
				{
					i:            0,
					j:            1,
					isPalindrome: false,
				},
				{
					i:            0,
					j:            2,
					isPalindrome: true,
				},
				{
					i:            1,
					j:            9,
					isPalindrome: true,
				},
				{
					i:            1,
					j:            6,
					isPalindrome: false,
				},
				{
					i:            1,
					j:            2,
					isPalindrome: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				this.memorize(arg.i, arg.j, arg.isPalindrome)
			}

			if !reflect.DeepEqual(want, this.memo) {
				t.Errorf("dp.memorize() = %v, want %v", this.memo, want)
			}
		})
	}
}

func Test_dp_isPalindrome(t *testing.T) {
	memo := map[int]map[int]bool{
		0: {
			0: true,
			1: false,
			2: true,
		},
		1: {
			9: true,
			6: false,
			2: false,
		},
	}
	type args struct {
		s string
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 01",
			args: args{
				s: "a",
				i: 0,
				j: 0,
			},
			want: true,
		},
		{
			name: "test 02",
			args: args{
				s: "aa",
				i: 0,
				j: 1,
			},
			want: false,
		},
		{
			name: "test 03",
			args: args{
				s: "aabaa",
				i: 0,
				j: 4,
			},
			want: true,
		},
		{
			name: "test 04",
			args: args{
				s: "aabaa",
				i: 1,
				j: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &dp{
				memo: memo,
			}
			if got := this.isPalindrome(tt.args.s, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("dp.isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
