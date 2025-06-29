package utils

func AddVector(src, toAdd *Vector) {
	src.X += toAdd.X
	src.Y += toAdd.Y
}

func AddMultipliedVector(src, toAdd *Vector, mul int32) {
	src.X += toAdd.X * mul
	src.Y += toAdd.Y * mul
}
