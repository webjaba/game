package systems

import (
	"game/internal/game/entities"
	"game/internal/game/settings"
	"game/internal/game/utils"
)

func Move[T entities.Entity](e []*T) {

}

// k []bool - булевая маска для клавиш WASD, true - если нажато, false - иначе
func MovePlayer(p *entities.Player, k []bool) {

	p.Vector.X = 0
	p.Vector.Y = 0

	if !k[0] && !k[1] && !k[2] && !k[3] {
		return
	}

	if k[0] { // W
		utils.AddVector(&p.Vector, &Up)
	}
	if k[1] { // A
		utils.AddVector(&p.Vector, &Left)
	}
	if k[2] { // S
		utils.AddVector(&p.Vector, &Down)
	}
	if k[3] { // D
		utils.AddVector(&p.Vector, &Right)
	}

	p.Coords.AddMul(&p.Vector, settings.PlayerSpeed)
}
