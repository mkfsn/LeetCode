package reverse_string_ii

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "abcdefg", k: 2},
			want: "bacdfeg",
		},
		{
			args: args{s: "abcd", k: 2},
			want: "bacd",
		},
		{
			args: args{s: "a", k: 2},
			want: "a",
		},
		{
			args: args{s: "abcdefg", k: 3},
			want: "cbadefg",
		},
		{
			args: args{s: "abcd", k: 4},
			want: "dcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
