package entities

import (
	"game/internal/game/settings"
	"game/internal/game/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer() *Entity {
	i := ebiten.NewImage(64, 96)
	i.Fill(color.RGBA{R: 255, A: 255})
	return &Entity{
		Position:  &utils.VecFloat64{X: settings.WindowSizeX / 2, Y: settings.WindowSizeY / 2},
		Direction: &utils.VecFloat64{X: 0, Y: 0},
		Type:      Player,
		Image:     i,
		Speed:     settings.PlayerSpeed,
	}
}

func NewZombie() *Entity {
	i := ebiten.NewImage(64, 96)
	i.Fill(color.RGBA{G: 255, A: 255})
	return &Entity{
		Position:  &utils.VecFloat64{X: 50, Y: 50},
		Direction: &utils.VecFloat64{X: 0, Y: 0},
		Type:      Zombie,
		Image:     i,
		Speed:     settings.ZombieSpeed,
		Ai:        &AiComponent{State: Neutral},
	}
}
