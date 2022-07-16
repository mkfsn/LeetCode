package subdomain_visit_count

import (
	"reflect"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	tests := []struct {
		name      string
		cpdomains []string
		want      []string
	}{
		{
			cpdomains: []string{
				"9001 discuss.leetcode.com",
			},
			want: []string{
				"9001 com",
				"9001 leetcode.com",
				"9001 discuss.leetcode.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.cpdomains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}
