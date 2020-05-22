package algorithms

import (
	"math"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test - 0",
			args: args{
				nums1: []int{1},
				nums2: []int{2},
			},
			want: 1.5,
		},
		{
			name: "test - 1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "test - 2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "test - 3",
			args: args{
				nums1: []int{},
				nums2: []int{3, 4},
			},
			want: 3.5,
		},
		{
			name: "test - 4",
			args: args{
				nums1: []int{},
				nums2: []int{1, 3, 4},
			},
			want: 3,
		},
		{
			name: "test - 5",
			args: args{
				nums1: []int{1, 3, 4},
				nums2: []int{},
			},
			want: 3,
		},
		{
			name: "test - 6",
			args: args{
				nums1: []int{1, 4},
				nums2: []int{},
			},
			want: 2.5,
		},
		{
			name: "test - 7",
			args: args{
				nums1: []int{},
				nums2: []int{},
			},
			want: math.MinInt64,
		},
		{
			name: "test - 8",
			args: args{
				nums1: []int{1},
				nums2: []int{100, 101, 101, 101},
			},
			want: 101,
		},
		{
			name: "test - 9",
			args: args{
				nums1: []int{1, 3, 4, 5, 6},
				nums2: []int{100, 101},
			},
			want: 5,
		},
		{
			name: "test - 10",
			args: args{
				nums1: []int{1, 3, 5, 7, 9},
				nums2: []int{2, 4, 6, 8, 9},
			},
			want: 5.5,
		},
		{
			name: "test - 11",
			args: args{
				nums1: []int{1, 3, 5, 5, 5},
				nums2: []int{2, 4, 6, 8, 9},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOf(t *testing.T) {
	type args struct {
		lst             []int
		start_inclusive int
		end_exclusive   int
		target          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty lst",
			args: args{
				lst:             []int{},
				start_inclusive: 0,
				end_exclusive:   0,
				target:          1,
			},
			want: 0,
		},
		{
			name: "Append to the end",
			args: args{
				lst:             []int{1},
				start_inclusive: 0,
				end_exclusive:   1,
				target:          2,
			},
			want: 1,
		},
		{
			name: "Put in the beginning",
			args: args{
				lst:             []int{5},
				start_inclusive: 0,
				end_exclusive:   1,
				target:          2,
			},
			want: 0,
		},
		{
			name: "Put in the middle - 0",
			args: args{
				lst:             []int{1, 5},
				start_inclusive: 0,
				end_exclusive:   1,
				target:          2,
			},
			want: 1,
		},
		{
			name: "Put in the middle - 1",
			args: args{
				lst:             []int{1, 5, 5, 5, 5},
				start_inclusive: 0,
				end_exclusive:   5,
				target:          2,
			},
			want: 1,
		},
		{
			name: "Put in the middle - 2",
			args: args{
				lst:             []int{1, 5, 5, 5, 5},
				start_inclusive: 1,
				end_exclusive:   5,
				target:          2,
			},
			want: 1,
		},
		{
			name: "Put in the end",
			args: args{
				lst:             []int{1, 5, 5, 5, 5},
				start_inclusive: 1,
				end_exclusive:   5,
				target:          6,
			},
			want: 5,
		},
		{
			name: "Put in the middle - 3",
			args: args{
				lst:             []int{1, 2, 3, 4, 5},
				start_inclusive: 0,
				end_exclusive:   5,
				target:          3,
			},
			want: 2,
		},
		{
			name: "Put in the middle - 4",
			args: args{
				lst:             []int{1, 2, 3, 4, 5},
				start_inclusive: 3,
				end_exclusive:   5,
				target:          3,
			},
			want: 3,
		},
		{
			name: "Left most - 5",
			args: args{
				lst:             []int{1, 2, 3, 3, 3, 3, 3, 4, 5},
				start_inclusive: 0,
				end_exclusive:   9,
				target:          3,
			},
			want: 2,
		},
		{
			name: "Left most - 6",
			args: args{
				lst:             []int{1, 2, 3, 3, 3, 3, 3, 4, 5},
				start_inclusive: 2,
				end_exclusive:   9,
				target:          3,
			},
			want: 2,
		},
		{
			name: "Left most - 7",
			args: args{
				lst:             []int{1, 2, 3, 3, 3, 3, 3, 4, 5},
				start_inclusive: 5,
				end_exclusive:   9,
				target:          3,
			},
			want: 5,
		},
		{
			name: "Normal find",
			args: args{
				lst:             []int{1, 2, 3, 4, 5, 7, 9},
				start_inclusive: 0,
				end_exclusive:   7,
				target:          6,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOf(tt.args.lst, tt.args.start_inclusive, tt.args.end_exclusive, tt.args.target); got != tt.want {
				t.Errorf("indexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averate(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averate(tt.args.lst); got != tt.want {
				t.Errorf("averate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedian(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "empty",
			args: args{
				lst: []int{},
			},
			want: float64(math.MinInt64),
		},
		{
			name: "odd",
			args: args{
				lst: []int{1, 2, 4},
			},
			want: 2.0,
		},
		{
			name: "even",
			args: args{
				lst: []int{1, 2, 4, 7},
			},
			want: 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedian(tt.args.lst); got != tt.want {
				t.Errorf("findMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
