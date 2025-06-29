package main

import (
	"game/internal/game/entities"
	"game/internal/game/systems"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game реализует интерфейс ebiten.Game
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
	// Отрисовка персонажа (пока просто красный квадрат)
	player := ebiten.NewImage(16, 16)
	player.Fill(color.RGBA{R: 255, A: 255}) // Красный цвет

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.Coords.X), float64(g.player.Coords.Y))
	screen.DrawImage(player, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // Фиксированный размер окна (можно масштабировать)
}

func main() {
	ebiten.SetWindowSize(640, 480) // Увеличим окно для удобства
	ebiten.SetWindowTitle("My Pixel Game")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
