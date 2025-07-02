package world

import (
	"game/internal/game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Room struct {
	Tiles  [][]int8
	Img    *ebiten.Image
	Coords *utils.Vector
}

func SetUpRoom() Room {
	r := Room{
		Tiles:  make([][]int8, 40),
		Coords: &utils.Vector{X: 0, Y: 0},
	}

	for i := range 40 {
		r.Tiles[i] = make([]int8, 40)
		for j := range 40 {
			r.Tiles[i][j] = StoneFloor
		}
	}

	r.Img = ebiten.NewImage(40*64, 40*64)

	for i, line := range r.Tiles {
		for j, tile := range line {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*64), float64(j*64))
			r.Img.DrawImage(
				Textures[tile],
				op,
			)
		}
	}

	return r
}
