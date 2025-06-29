package game

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"game/internal/game/systems"
	"game/internal/game/utils"
	"game/internal/game/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player      entities.Player
	fullscreen  bool
	room        world.Room
	roomTexture *ebiten.Image
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.fullscreen = !g.fullscreen
		ebiten.SetFullscreen(g.fullscreen)
	}

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
	systems.DrawRoom(screen, g.roomTexture)
	systems.DrawPlayer(screen, &g.player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.WindowSizeX, settings.WindowSizeY
}

func SetUp() *Game {
	ebiten.SetWindowSize(settings.WindowSizeX, settings.WindowSizeY)

	i := ebiten.NewImage(64, 64)
	i.Fill(color.RGBA{R: 255, A: 255})

	// тест отрисовки комнаты
	r := world.SetUpRoom()

	staticLayer := ebiten.NewImage(10*64, 10*64)

	for i, line := range r {
		for j, tile := range line {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*64), float64(j*64))
			staticLayer.DrawImage(
				world.Textures[tile],
				op,
			)
		}
	}
	// конец теста

	return &Game{
		player: entities.Player{
			Coords: utils.Vector{X: settings.WindowSizeX / 2, Y: settings.WindowSizeY / 2},
			Vector: utils.Vector{},
			Img:    i,
		},
		room:        r,
		roomTexture: staticLayer,
	}
}
