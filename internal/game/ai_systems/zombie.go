package aisystems

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"math"
)

func UpdateZombie(player, e *entities.Entity) {
	dx, dy, l := updateState(player, e)

	if e.Ai.State == entities.Aggressive {
		if l != 0 {
			e.Direction.X = dx / l
			e.Direction.Y = dy / l
			return
		}
		e.Direction.X = dx
		e.Direction.Y = dy
	}
}

func updateState(p, e *entities.Entity) (dx, dy, l float64) {
	dx = p.Position.X - e.Position.X
	dy = p.Position.Y - e.Position.Y

	l = math.Hypot(dx, dy)

	if l <= settings.ZombieActivationRange {
		e.Ai.State = entities.Aggressive
		return
	}

	e.Direction.X = 0
	e.Direction.Y = 0

	if l > settings.ZombieActivationRange && l < settings.MaxDistance {
		e.Ai.State = entities.Neutral
		return
	}
	e.Ai.State = entities.Stopped
	return
}
