package internal

// Velocity represents a velocity in a 2D grid.
type Velocity struct {
	dx, dy int
}

// NewVelocity checks if the velocity is within range.
func (v Velocity) IsWithinRange() bool {
	const maxVelocity = 3
	const minVelocity = -3

	return v.dx >= minVelocity && v.dx <= maxVelocity && v.dy >= minVelocity && v.dy <= maxVelocity
}
