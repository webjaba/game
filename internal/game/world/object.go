package world

import (
	"game/internal/game/settings"
	"game/internal/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Empty int8 = iota
	StoneFloor
)

type Room [][]int8

var (
	Textures = map[int8]*ebiten.Image{
		StoneFloor: utils.LoadTexture(settings.StoneFloorTexture),
	}
)
