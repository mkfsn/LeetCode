package greatest_common_divisor_of_strings

import (
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			str1: "AAAA",
			str2: "A",
			want: "A",
		},
		{
			str1: "A",
			str2: "AAAA",
			want: "A",
		},
		{
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			str1: "AB",
			str2: "AA",
			want: "",
		},
		{
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
