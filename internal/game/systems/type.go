package systems

import (
	"game/internal/game/entities"
	"game/internal/game/utils"
)

/*
координаты отсчитываются от левой верхней точки экрана
поэтому чем ниже позиция, тем больше Y, соответственно
пространство инвертировано по горизонтальной оси
*/
var (
	Up    = utils.VecFloat64{X: 0, Y: -1}
	Down  = utils.VecFloat64{X: 0, Y: 1}
	Left  = utils.VecFloat64{X: -1, Y: 0}
	Right = utils.VecFloat64{X: 1, Y: 0}
)

type System interface {
	Update([]*entities.Entity)
}
