package entities

import (
	"game/internal/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Coords utils.Vector
	Vector utils.Vector
	Img    *ebiten.Image
}
