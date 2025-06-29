package game

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"game/internal/game/systems"
	"game/internal/game/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player entities.Player
}

func (g *Game) Update() error {

	wasd := []bool{
		ebiten.IsKeyPressed(ebiten.KeyW),
		ebiten.IsKeyPressed(ebiten.KeyA),
		ebiten.IsKeyPressed(ebiten.KeyS),
		ebiten.IsKeyPressed(ebiten.KeyD),
	}
	systems.MovePlayer(&g.player, wasd)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	systems.DrawPlayer(screen, &g.player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.WindowSizeX, settings.WindowSizeY
}

func SetUp() *Game {
	ebiten.SetWindowSize(settings.WindowSizeX, settings.WindowSizeY)

	i := ebiten.NewImage(64, 64)
	i.Fill(color.RGBA{R: 255, A: 255})

	return &Game{
		player: entities.Player{
			Coords: utils.Vector{X: settings.WindowSizeX / 2, Y: settings.WindowSizeY / 2},
			Vector: utils.Vector{},
			Img:    i,
		},
	}
}
