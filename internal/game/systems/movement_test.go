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
		in   []*entities.Entity
		want []*entities.Entity
	}{
		{
			name: "nil",
			in:   []*entities.Entity{{}},
			want: []*entities.Entity{{}},
		},
		{
			name: "dont move",
			in:   []*entities.Entity{mustBuildEntity(nil, nil)},
			want: []*entities.Entity{mustBuildEntity(nil, nil)},
		},
		{
			name: "move one",
			in:   []*entities.Entity{mustBuildEntity(nil, &utils.Vector{X: 1, Y: -1})},
			want: []*entities.Entity{mustBuildEntity(
				&utils.Vector{X: settings.PlayerSpeed, Y: -settings.PlayerSpeed},
				&utils.Vector{X: 1, Y: -1},
			)},
		},
		{
			name: "move many",
			in: []*entities.Entity{
				mustBuildEntity(nil, &utils.Vector{X: 1, Y: -1}),
				mustBuildEntity(&utils.Vector{X: 10, Y: 10}, &utils.Vector{X: 1, Y: -1}),
			},
			want: []*entities.Entity{
				mustBuildEntity(
					&utils.Vector{X: settings.PlayerSpeed, Y: -settings.PlayerSpeed},
					&utils.Vector{X: 1, Y: -1},
				),
				mustBuildEntity(
					&utils.Vector{X: 10 + settings.PlayerSpeed, Y: 10 - settings.PlayerSpeed},
					&utils.Vector{X: 1, Y: -1},
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &MovementSystem{}
			s.Update(tt.in)

			for i, w := range tt.want {
				require.Equal(t, w.Position, tt.in[i].Position, "Position")
				require.Equal(t, w.Direction, tt.in[i].Direction, "Direction")
			}
		})
	}
}

func mustBuildEntity(Position, Direction *utils.Vector) *entities.Entity {
	p := &entities.Entity{
		Speed:     settings.PlayerSpeed,
		Position:  &utils.Vector{X: 0, Y: 0},
		Direction: &utils.Vector{X: 0, Y: 0},
	}
	if Position != nil {
		p.Position = Position
	}
	if Direction != nil {
		p.Direction = Direction
	}
	return p
}
