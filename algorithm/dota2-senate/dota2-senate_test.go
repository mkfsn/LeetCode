package dota2_senate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_predictPartyVictory(t *testing.T) {
	tests := []struct {
		senate string
		want   string
	}{
		{
			senate: "RD",
			want:   "Radiant",
		},
		{
			senate: "RDD",
			want:   "Dire",
		},
		{
			senate: "DDD",
			want:   "Dire",
		},
		{
			senate: "DRDRDRDRDRDRDRR",
			want:   "Dire",
		},
		{
			senate: "DRDRDRDRDRDRDRRD",
			want:   "Dire",
		},
		{
			senate: "RRDDD",
			want:   "Radiant",
		},
		{
			senate: "DRRDRDRDRDDRDRDR",
			// DxRxRxRxRxDxDxDx
			// DRRRRDDD
			// DxRRRxxx
			// DRRR
			// xxRR
			want: "Radiant",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := predictPartyVictory(tt.senate)
			assert.Equal(t, tt.want, got)
		})
	}
}
