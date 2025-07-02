package entities

import (
	"game/internal/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type EffectType uint8

type Duration uint8

type EntityType uint8

const (
	Burning EffectType = iota
	Freezing
	Stunning
)

const (
	Player EntityType = iota
	Zombie
)

type Entity struct {
	Position  *utils.Vector
	Direction *utils.Vector
	Speed     int32 // Возможна оптимизация
	Effects   map[EffectType]Duration
	Image     *ebiten.Image
	Type      EntityType
}
