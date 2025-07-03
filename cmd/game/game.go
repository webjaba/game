package game

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"game/internal/game/systems"
	"game/internal/game/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player   *entities.Entity
	entities []*entities.Entity
	systems  []systems.System
	// camera     world.Camera
	renderSystem     systems.RenderSystem
	handleKeysSystem systems.HandleKeysSystem
	fullscreen       bool
	room             world.Room
}

func (g *Game) Update() error {
	g.handleKeysSystem.Update(g.player)
	for _, s := range g.systems {
		s.Update(g.entities)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderSystem.Render(screen, g.entities)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.WindowSizeX, settings.WindowSizeY
}

func New() *Game {
	ebiten.SetWindowSize(settings.WindowSizeX, settings.WindowSizeY)

	r := world.SetUpRoom()

	p := entities.NewPlayer()

	return &Game{
		renderSystem:     *systems.NewRenderSystem(p, r.Img),
		handleKeysSystem: systems.HandleKeysSystem{},
		player:           p,
		room:             r,
		entities: []*entities.Entity{
			p,
			entities.NewZombie(),
		},
		systems: []systems.System{
			&systems.MovementSystem{},
			systems.NewAiSystem(p),
		},
	}
}
