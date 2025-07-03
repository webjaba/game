package entities

import (
	"game/internal/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type AiComponent struct {
	State AiComponentState
}

type Entity struct {
	Position  *utils.VecFloat64
	Direction *utils.VecFloat64
	Speed     float64
	Effects   map[EffectType]Duration
	Image     *ebiten.Image
	Type      EntityType
	Ai        *AiComponent
}
