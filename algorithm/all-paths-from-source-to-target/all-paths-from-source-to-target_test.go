package all_paths_from_source_to_target

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{
			graph: [][]int{
				{1, 2},
				{3},
				{3},
				{},
			},
			want: [][]int{
				{0, 1, 3},
				{0, 2, 3},
			},
		},
		{
			graph: [][]int{
				{4, 3, 1},
				{3, 2, 4},
				{3},
				{4},
				{},
			},
			want: [][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			},
		},
		{
			graph: [][]int{
				{1, 2},
				{},
				{},
				{},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
