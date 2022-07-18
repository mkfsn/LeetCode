package find_if_path_exists_in_graph

import (
	"testing"
)

func Test_validPath(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{
			n: 10,
			edges: [][]int{
				{0, 7},
				{0, 8},
				{6, 1},
				{2, 0},
				{0, 4},
				{5, 8},
				{4, 7},
				{1, 3},
				{3, 5},
				{6, 5},
			},
			source:      7,
			destination: 5,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPath(tt.n, tt.edges, tt.source, tt.destination); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
