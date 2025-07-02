package entities

import (
	"game/internal/game/settings"
	"game/internal/game/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer() *Entity {
	i := ebiten.NewImage(64, 64)
	i.Fill(color.RGBA{R: 255, A: 255})
	return &Entity{
		Position:  &utils.Vector{X: settings.WindowSizeX / 2, Y: settings.WindowSizeY / 2},
		Direction: &utils.Vector{X: 0, Y: 0},
		Type:      Player,
		Image:     i,
		Speed:     settings.PlayerSpeed,
	}
}
