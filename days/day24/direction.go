package day24

// All possible directions.
const (
	East      = "e"
	NorthEast = "ne"
	SouthEast = "se"
	West      = "w"
	NorthWest = "nw"
	SouthWest = "sw"
)

// Magic numbers.
const (
	offsetIncementSingleDigit = 1
	offsetIncementultiDigit   = 2
)

// Direction is direction.
type Direction string

func directionFromString(current string) (Direction, int, error) {
	if current[0] == 'w' {
		return West, offsetIncementSingleDigit, nil
	}

	if current[0] == 'e' {
		return East, offsetIncementSingleDigit, nil
	}

	switch current {
	case "nw":
		return NorthWest, offsetIncementultiDigit, nil
	case "sw":
		return SouthWest, offsetIncementultiDigit, nil
	case "ne":
		return NorthEast, offsetIncementultiDigit, nil
	case "se":
		return SouthEast, offsetIncementultiDigit, nil
	}

	return "", 0, UnknownDirection{Direction: current}
}
