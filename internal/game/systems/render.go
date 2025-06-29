package systems

import (
	"game/internal/game/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPlayer(screen *ebiten.Image, p *entities.Player) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Coords.X), float64(p.Coords.Y))
	screen.DrawImage(p.Img, op)
}
