package systems

import (
	"game/internal/game/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type HandleKeysSystem struct{}

func (s *HandleKeysSystem) Update(player *entities.Entity) {
	player.Direction.X = 0
	player.Direction.Y = 0

	if ebiten.IsKeyPressed(ebiten.KeyW) { // W
		player.Direction.Add(&Up)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) { // A
		player.Direction.Add(&Left)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) { // S
		player.Direction.Add(&Down)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) { // D
		player.Direction.Add(&Right)
	}

	player.Position.AddMul(player.Direction, player.Speed)
}
