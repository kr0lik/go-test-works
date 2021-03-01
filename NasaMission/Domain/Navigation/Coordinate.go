package Navigation

type Coordinate struct {
	x, y int
}

func CreateCoordinate(x int, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

func (coordinate *Coordinate) GetX() int {
	return coordinate.x
}

func (coordinate *Coordinate) GetY() int {
	return coordinate.y
}
