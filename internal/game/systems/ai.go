package systems

import (
	aisystems "game/internal/game/ai_systems"
	"game/internal/game/entities"
)

type AiSystem struct {
	player     *entities.Entity
	entitiesAi map[entities.EntityType]func(player, entity *entities.Entity)
}

func (s *AiSystem) Update(ents []*entities.Entity) {
	for _, e := range ents {
		if e.Ai != nil {
			s.entitiesAi[e.Type](s.player, e)
		}
	}
}

func NewAiSystem(p *entities.Entity) *AiSystem {
	return &AiSystem{
		player: p,
		entitiesAi: map[entities.EntityType]func(player *entities.Entity, entity *entities.Entity){
			entities.Zombie: aisystems.UpdateZombie,
		},
	}
}
