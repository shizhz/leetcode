package algorithms

import (
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
