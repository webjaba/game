package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddVector(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		src, toAdd Vector
		want       Vector
	}{
		{
			name:  "zero vector",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: 0, Y: 0},
			want:  Vector{X: 0, Y: 0},
		},
		{
			name:  "add positive",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: 1, Y: 1},
			want:  Vector{X: 1, Y: 1},
		},
		{
			name:  "add negative",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: -1, Y: -1},
			want:  Vector{X: -1, Y: -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			AddVector(&tt.src, &tt.toAdd)

			require.Equal(t, tt.src.X, tt.want.X, "X")
			require.Equal(t, tt.src.Y, tt.want.Y, "Y")
		})
	}
}

func TestAddMultipliedVector(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		src, toAdd Vector
		mul        int32
		want       Vector
	}{
		{
			name:  "zero vector",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: 0, Y: 0},
			want:  Vector{X: 0, Y: 0},
		},
		{
			name:  "add positive",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: 1, Y: 1},
			want:  Vector{X: 2, Y: 2},
			mul:   2,
		},
		{
			name:  "add negative",
			src:   Vector{X: 0, Y: 0},
			toAdd: Vector{X: -1, Y: -1},
			want:  Vector{X: -2, Y: -2},
			mul:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			AddMultipliedVector(&tt.src, &tt.toAdd, tt.mul)

			require.Equal(t, tt.want.X, tt.src.X, "X")
			require.Equal(t, tt.want.Y, tt.src.Y, "Y")
		})
	}
}
