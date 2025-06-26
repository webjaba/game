package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game реализует интерфейс ebiten.Game
type Game struct {
	playerX, playerY float64
}

func (g *Game) Update() error {
	// Движение персонажа (WASD или стрелки)
	speed := 3.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.playerX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.playerX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.playerY -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.playerY += speed
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Отрисовка персонажа (пока просто красный квадрат)
	player := ebiten.NewImage(16, 16)
	player.Fill(color.RGBA{R: 255, A: 255}) // Красный цвет

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
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
