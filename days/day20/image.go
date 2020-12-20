package day20

// Image is an image.
type Image struct {
	Corner       []*Tile
	Side         []*Tile
	Middle       []*Tile
	NumberOfRows int
}

// NewImage creates a new Image.
func NewImage() *Image {
	return &Image{
		Corner:       []*Tile{},
		Side:         []*Tile{},
		Middle:       []*Tile{},
		NumberOfRows: 0,
	}
}

// AddCorner saves a corner piece.
func (image *Image) AddCorner(tile *Tile) {
	image.Corner = append(image.Corner, tile)
}

// AddSide saves a side piece.
func (image *Image) AddSide(tile *Tile) {
	image.Side = append(image.Side, tile)
}

// AddMiddle saves a middle piece.
func (image *Image) AddMiddle(tile *Tile) {
	image.Middle = append(image.Middle, tile)
}
