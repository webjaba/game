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

func (v *Vector) Normalize() {
	if v.X < 0 {
		v.X /= -v.X
	} else if v.X > 0 {
		v.X /= v.X
	}
	if v.Y < 0 {
		v.Y /= -v.Y
	} else if v.Y > 0 {
		v.Y /= v.Y
	}
}
