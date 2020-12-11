package day11

// Available positions.
const (
	Empty    = 'L'
	Occupied = '#'
	Floor    = '.'
)

// Position represents a spot on the ferry.
// It can be Empty, Occupied or Floor.
type Position byte

// Row is a series of Positions.
type Row []Position

// Ferry is the ferry itself with rows.
type Ferry struct {
	Rows []Row
}

// NewFerry creates a new Ferry.
func NewFerry() Ferry {
	return Ferry{
		Rows: []Row{},
	}
}

// AddRow adds a row onto the Ferry.
func (f *Ferry) AddRow(row Row) {
	f.Rows = append(f.Rows, row)
}

// ParseRow parses a string repesentation of a row
// into a Row.
func ParseRow(input string) (Row, error) {
	row := []Position{}

	for _, ch := range input {
		switch ch {
		case Empty:
			row = append(row, Empty)
		case Occupied:
			row = append(row, Occupied)
		case Floor:
			row = append(row, Floor)
		default:
			return row, UnkownCharacter{Character: ch}
		}
	}

	return row, nil
}

// Count how many spots the Ferry has from a given Position.
func (f *Ferry) Count(position Position) int {
	count := 0

	for _, row := range f.Rows {
		for _, p := range row {
			if p == position {
				count++
			}
		}
	}

	return count
}

// GetPosition returns with the position on a given spot
// and a boolean if it's off the grid or not.
// If a given (x,y) coordinate is off the grid, it will
// be always floor and the boolean true.
func (f *Ferry) GetPosition(x, y int) (Position, bool) {
	if y < 0 || y >= len(f.Rows) || x < 0 || x >= len(f.Rows[y]) {
		return Floor, true
	}

	return f.Rows[y][x], false
}

// Adjacents collects and return with all adjacent positions.
// If long is true, it goes "until they see a seat",
// so it will ignore Floor type.
func (f *Ferry) Adjacents(x, y int, long bool) []Position {
	adjacents := []Position{}

	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			if dx|dy == 0 {
				continue
			}

			var (
				pos        Position
				offTheGrid bool
				counter    int = 1
			)

			for {
				pos, offTheGrid = f.GetPosition(x+(dx*counter), y+(dy*counter))

				if !long || pos != Floor || offTheGrid {
					break
				}

				counter++
			}

			adjacents = append(adjacents, pos)
		}
	}

	return adjacents
}

// CountAdjacents counts all adjacent positions for a given value.
// If long is true, it goes "until they see a seat",
// so it will ignore Floor type.
func (f *Ferry) CountAdjacents(x, y int, long bool, position Position) (count int) {
	adjacents := f.Adjacents(x, y, long)

	for _, v := range adjacents {
		if v == position {
			count++
		}
	}

	return
}

// Cycle is a simple life cycle.
func (f *Ferry) Cycle(occupationThreshold int, long bool) (changes int) {
	updated := []Row{}

	for y, row := range f.Rows {
		updatedRow := []Position{}

		for x, pos := range row {
			newPos := pos

			switch pos {
			case Occupied:
				if f.CountAdjacents(x, y, long, Occupied) > occupationThreshold {
					newPos = Empty

					changes++
				}
			case Empty:
				if f.CountAdjacents(x, y, long, Occupied) < 1 {
					newPos = Occupied

					changes++
				}
			}

			updatedRow = append(updatedRow, newPos)
		}

		updated = append(updated, updatedRow)
	}

	f.Rows = updated

	return changes
}
