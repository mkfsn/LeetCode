package decode_string

import (
	"testing"
)

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			s:    "2[a]3[c4[d]]e",
			want: "aacddddcddddcdddde",
		},
		{
			s: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			// zzz(2(yypq4(jkjkef))ef
			// zzz(2(yypqjkjkefjkjkefjkjkefjkjkef))ef
			//     zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef
			want: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
			//    "zzzyypqjkjkefefjkjkefefjkjkefefjkjkefefyypqjkjkefefjkjkefefjkjkefefjkjkefef"
			//                  |
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
