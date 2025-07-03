package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVecFloat64_AddMul(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		src, toAdd VecFloat64
		mul        float64
		want       VecFloat64
	}{
		{
			name:  "zero VecFloat64",
			src:   VecFloat64{X: 0, Y: 0},
			toAdd: VecFloat64{X: 0, Y: 0},
			want:  VecFloat64{X: 0, Y: 0},
		},
		{
			name:  "add positive",
			src:   VecFloat64{X: 0, Y: 0},
			toAdd: VecFloat64{X: 1, Y: 1},
			want:  VecFloat64{X: 2, Y: 2},
			mul:   2,
		},
		{
			name:  "add negative",
			src:   VecFloat64{X: 0, Y: 0},
			toAdd: VecFloat64{X: -1, Y: -1},
			want:  VecFloat64{X: -2, Y: -2},
			mul:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.src.AddMul(&tt.toAdd, tt.mul)

			require.Equal(t, tt.want.X, tt.src.X, "X")
			require.Equal(t, tt.want.Y, tt.src.Y, "Y")
		})
	}
}
