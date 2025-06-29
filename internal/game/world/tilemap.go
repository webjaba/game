package world

func SetUpRoom() Room {
	r := make(Room, 10)

	for i := range 10 {
		r[i] = make([]int8, 10)
		for j := range 10 {
			r[i][j] = StoneFloor
		}
	}

	return r
}
