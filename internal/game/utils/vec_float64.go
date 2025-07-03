package utils

type VecFloat64 struct {
	X float64
	Y float64
}

func (v *VecFloat64) AddMul(toAdd *VecFloat64, mul float64) {
	v.X += toAdd.X * mul
	v.Y += toAdd.Y * mul
}

func (v *VecFloat64) Add(toAdd *VecFloat64) {
	v.X += toAdd.X
	v.Y += toAdd.Y
}
