package entities

type EffectType uint8

type Duration uint8

type EntityType uint8

type AiComponentState uint8

const (
	Burning EffectType = iota
	Freezing
	Stunning
)

const (
	Player EntityType = iota
	Zombie
)

const (
	Stopped AiComponentState = iota
	Neutral
	Aggressive
)
