package systems

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"game/internal/game/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMovePlayer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    entities.Player
		wasd []bool
		want entities.Player
	}{
		{
			name: "no key pressed",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			wasd: []bool{false, false, false, false},
		},
		{
			name: "W",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 0, Y: -settings.PlayerSpeed},
				Vector: utils.Vector{X: 0, Y: -1},
			},
			wasd: []bool{true, false, false, false},
		},
		{
			name: "A",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: -settings.PlayerSpeed, Y: 0},
				Vector: utils.Vector{X: -1, Y: 0},
			},
			wasd: []bool{false, true, false, false},
		},
		{
			name: "S",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 0, Y: settings.PlayerSpeed},
				Vector: utils.Vector{X: 0, Y: 1},
			},
			wasd: []bool{false, false, true, false},
		},
		{
			name: "D",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: settings.PlayerSpeed, Y: 0},
				Vector: utils.Vector{X: 1, Y: 0},
			},
			wasd: []bool{false, false, false, true},
		},
		{
			name: "WASD",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			wasd: []bool{true, true, true, true},
		},
		{
			name: "WAD",
			p: entities.Player{
				Coords: utils.Vector{X: 0, Y: 0},
				Vector: utils.Vector{X: 0, Y: 0},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 0, Y: -settings.PlayerSpeed},
				Vector: utils.Vector{X: 0, Y: -1},
			},
			wasd: []bool{true, true, false, true},
		},
		{
			name: "WA->WAD",
			p: entities.Player{
				Coords: utils.Vector{X: 10, Y: 10},
				Vector: utils.Vector{X: -1, Y: -1},
			},
			want: entities.Player{
				Coords: utils.Vector{X: 10, Y: 10 - settings.PlayerSpeed},
				Vector: utils.Vector{X: 0, Y: -1},
			},
			wasd: []bool{true, true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			MovePlayer(&tt.p, tt.wasd)

			require.Equal(t, tt.want.Coords.X, tt.p.Coords.X, "Coords.X")
			require.Equal(t, tt.want.Coords.Y, tt.p.Coords.Y, "Coords.Y")
			require.Equal(t, tt.want.Vector.Y, tt.p.Vector.Y, "Vector.X")
			require.Equal(t, tt.want.Vector.Y, tt.p.Vector.Y, "Vector.Y")
		})
	}
}
