package day03

type vector2D struct {
	X int64
	Y int64
}

type snowField struct {
	Position vector2D
	Area     []string
	Length   int64
	Height   int64
}

func newMap(area []string) snowField {
	return snowField{
		Position: vector2D{X: 0, Y: 0},
		Area:     area,
		Length:   int64(len(area[0])),
		Height:   int64(len(area)),
	}
}

// Grow is a lie, it's just a simple math trick.
func (m *snowField) Grow() {
	if m.Position.X >= m.Length {
		m.Position.X -= m.Length
	}
}

func (m *snowField) Step(vec vector2D) bool {
	m.Position.X += vec.X
	m.Position.Y += vec.Y

	m.Grow()

	return m.Position.Y < m.Height
}

func (m *snowField) CurrentSquare() byte {
	return m.Area[m.Position.Y][m.Position.X]
}

func (m *snowField) Reset() {
	m.Position = vector2D{X: 0, Y: 0}
}
