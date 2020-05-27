package algorithms

import (
	"testing"
)

func Test_buildPattern(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test 01",
			args: "",
			want: "",
		},
		{
			name: "test 02",
			args: "a",
			want: "a",
		},
		{
			name: "test 03",
			args: "a.*",
			want: "a.*",
		},
		{
			name: "test 04",
			args: "a.**",
			want: "a.*",
		},
		{
			name: "test 05",
			args: "a..*",
			want: "a..*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPattern(tt.args); got != tt.want {
				t.Errorf("buildPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_match(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 01",
			args: args{
				s: "",
				p: "",
			},
			want: true,
		},
		{
			name: "test 02",
			args: args{
				s: "",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test 03",
			args: args{
				s: "a",
				p: "",
			},
			want: false,
		},
		{
			name: "test 04",
			args: args{
				s: "a",
				p: ".",
			},
			want: true,
		},
		{
			name: "test 05",
			args: args{
				s: "a",
				p: "",
			},
			want: false,
		},
		{
			name: "test 06",
			args: args{
				s: "a",
				p: "a",
			},
			want: true,
		},
		{
			name: "test 07",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "test 08",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "test 09",
			args: args{
				s: "aa",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test 10",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "test 11",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "test 12",
			args: args{
				s: "a",
				p: "ab*",
			},
			want: true,
		},
		{
			name: "test 13",
			args: args{
				s: "",
				p: "c*c*c***",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
