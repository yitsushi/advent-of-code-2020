package day11

import "fmt"

// UnkownCharacter occurs when the parser got an invalid characters.
// Valid characters are:
//   L -> empty
//   # -> occupied
//   . -> floor
type UnkownCharacter struct {
	Character rune
}

func (e UnkownCharacter) Error() string {
	return fmt.Sprintf("Unknown character: %c", e.Character)
}
