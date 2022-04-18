package fibonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n: 2, want: 1,
		},
		{
			n: 3, want: 2,
		},
		{
			n: 4, want: 3,
		},
		{
			n: 30, want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := fib(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
