package minimum_sum_of_four_digit_number_after_splitting_digits

import (
	"testing"
)

func Test_minimumSum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{num: 2932},
			want: 52,
		},
		{
			args: args{num: 4009},
			want: 13,
		},
		{
			args: args{num: 5787},
			want: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.num); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
