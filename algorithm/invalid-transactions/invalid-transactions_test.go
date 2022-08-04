package invalid_transactions

import (
	"reflect"
	"testing"
)

func Test_invalidTransactions(t *testing.T) {
	tests := []struct {
		name         string
		transactions []string
		want         []string
	}{
		{
			transactions: []string{
				"bob,689,1910,barcelona", // this > 1000
				"alex,696,122,bangkok",
				"bob,832,1726,barcelona", // this, > 1000
				"bob,820,596,bangkok",    // this
				"chalicefy,217,669,barcelona",
				"bob,175,221,amsterdam",
			},
			// output: ["bob,689,1910,barcelona","bob,832,1726,barcelona"]
			want: []string{
				"bob,689,1910,barcelona",
				"bob,820,596,bangkok",
				"bob,832,1726,barcelona",
			},
		},
		{
			transactions: []string{
				"alice,20,800,mtv", "alice,50,1200,mtv",
			},
			want: []string{
				"alice,50,1200,mtv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invalidTransactions(tt.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invalidTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}
