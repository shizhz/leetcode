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
