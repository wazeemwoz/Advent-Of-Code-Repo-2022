package types

type Coord struct {
	X int
	Y int
}

func (coord Coord) Add(other Coord) Coord {
	return Coord{coord.X + other.X, coord.Y + other.Y}
}
