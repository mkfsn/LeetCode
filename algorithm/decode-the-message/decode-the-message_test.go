package decode_the_message

import (
	"testing"
)

func Test_decodeMessage(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		message string
		want    string
	}{
		{
			key:     "the quick brown fox jumps over the lazy dog",
			message: "vkbs bs t suepuv",
			want:    "this is a secret",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeMessage(tt.key, tt.message); got != tt.want {
				t.Errorf("decodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
