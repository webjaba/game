package systems

import (
	"game/internal/game/entities"
)

type MovementSystem struct{}

func (m *MovementSystem) Update(ents []*entities.Entity) {
	for _, e := range ents {
		if e.Direction != nil && e.Position != nil {
			e.Position.AddMul(e.Direction, e.Speed)
		}
	}
}
