package systems

import (
	"game/internal/game/entities"
	"game/internal/game/settings"

	"github.com/hajimehoshi/ebiten/v2"
)

type RenderSystem struct {
	player     *entities.Entity
	background *ebiten.Image
}

func (s *RenderSystem) Render(screen *ebiten.Image, ents []*entities.Entity) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(-s.player.Position.X+settings.WindowSizeX/2-32),
		float64(-s.player.Position.Y+settings.WindowSizeY/2-32),
	)
	screen.DrawImage(s.background, op)
	for _, e := range ents {
		if e.Image != nil && e.Position != nil {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(e.Position.X-s.player.Position.X+settings.WindowSizeX/2-32),
				float64(e.Position.Y-s.player.Position.Y+settings.WindowSizeY/2-32),
			)
			screen.DrawImage(e.Image, op)
		}
	}
}

func (s *RenderSystem) UpdateBackground(b *ebiten.Image) {
	s.background = b
}

func NewRenderSystem(p *entities.Entity, b *ebiten.Image) *RenderSystem {
	return &RenderSystem{
		player:     p,
		background: b,
	}
}
