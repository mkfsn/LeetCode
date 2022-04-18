package latest_time_by_replacing_hidden_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumTime(t *testing.T) {
	tests := []struct {
		time string
		want string
	}{
		{
			time: "2?:?0",
			want: "23:50",
		},
		{
			time: "0?:3?",
			want: "09:39",
		},
		{
			time: "1?:22",
			want: "19:22",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maximumTime(tt.time)
			assert.Equal(t, tt.want, got)
		})
	}
}
