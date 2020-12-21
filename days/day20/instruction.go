package day20

// FitInstruction is all the instructions to fit next to the
// the other tile.
type FitInstruction struct {
	Rotation int
	FlipX    bool
	FlipY    bool
}

// Execute instructions on a tile.
func (f FitInstruction) Execute(tile *Tile) {
	if f.Rotation > 0 {
		tile.Rotate(f.Rotation)
	}

	if f.FlipX {
		tile.FlipX()
	}

	if f.FlipY {
		tile.FlipY()
	}
}
