package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	fmt.Println("test 0")
	type args struct {
		nums1 []int
		nums2 []int
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
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_medianIndices(t *testing.T) {
	type args struct {
		tom   []int
		jerry []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Both empty",
			args: args{
				tom:   []int{},
				jerry: []int{},
			},
			want: []int{},
		},
		{
			name: "Both odd",
			args: args{
				tom:   []int{0},
				jerry: []int{1},
			},
			want: []int{0, 1},
		},
		{
			name: "Jerry empty",
			args: args{
				tom:   []int{0, 1, 2},
				jerry: []int{},
			},
			want: []int{1},
		},
		{
			name: "Both not empty",
			args: args{
				tom:   []int{0, 1, 2, 3},
				jerry: []int{4, 5},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianIndices(tt.args.tom, tt.args.jerry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianIndices() = %v, want %v", got, tt.want)
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
