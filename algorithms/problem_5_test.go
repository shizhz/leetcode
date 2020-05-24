package algorithms

import (
	"fmt"
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
			this := tt.this
			if got := this.isPalindrome(tt.args); got != tt.want {
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
