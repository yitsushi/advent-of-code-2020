package math

// OrintedObject2D is an object that has position and orientation.
type OrintedObject2D struct {
	Position Vector2D
	Facing   Vector2D
}

// Move the object based on its orientation (forward).
func (o *OrintedObject2D) Move(unit float64) {
	o.Position.X += unit * o.Facing.X
	o.Position.Y += unit * o.Facing.Y
}

// Rotate the object.
func (o *OrintedObject2D) Rotate(degree float64) {
	o.Facing.Rotate(degree)
}
