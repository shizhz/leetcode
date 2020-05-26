package algorithms

import (
	"fmt"
	"strings"
	"testing"
)

func TestRune(t *testing.T) {
	runes := []string{"a", "", "b", "c", "你"}

	fmt.Println(strings.Join(runes, ""))
}

func Test_zigzag(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 01",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "test 02",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "test 03",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 1,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "test 04",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 100,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "test 05",
			args: args{
				s:       "",
				numRows: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzag(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("zigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}
