package utils

type Vector struct {
	X int32
	Y int32
}

func (v *Vector) AddMul(toAdd *Vector, mul int32) {
	v.X += toAdd.X * mul
	v.Y += toAdd.Y * mul
}

func (v *Vector) Add(toAdd *Vector) {
	v.X += toAdd.X
	v.Y += toAdd.Y
}
