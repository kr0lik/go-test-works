package Enum

type Command string

const (
	Left  Command = "L"
	Right Command = "R"
	Move  Command = "M"
)

func (action Command) GetValues() []Command {
	return []Command{Left, Right, Move}
}

func (action Command) IsValid() bool {
	for _, check := range action.GetValues() {
		if check == action {
			return true
		}
	}

	return false
}
