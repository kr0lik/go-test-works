package Enum

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

func (direction Direction) GetValues() []Direction {
	return []Direction{North, East, South, West}
}

func (direction Direction) IsValid() bool {
	for _, check := range direction.GetValues() {
		if check == direction {
			return true
		}
	}

	return false
}
