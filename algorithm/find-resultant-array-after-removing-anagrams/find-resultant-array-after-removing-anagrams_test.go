package find_resultant_array_after_removing_anagrams

import (
	"reflect"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			want:  []string{"abba", "cd"},
		},

		{
			words: []string{"a", "b", "a"},
			want:  []string{"a", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAnagrams(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
