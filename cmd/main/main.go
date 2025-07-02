package main

import (
	"game/cmd/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("My Pixel Game")

	g := game.New()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
