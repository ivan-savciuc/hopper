package internal

import "fmt"

// Point represents a point in a 2D grid.
type Point struct {
	X, Y int
}

// NewPoint creates a new point with the given x and y coordinates.
func (p Point) IsWithinRange(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// NewPoint creates a new point with the given x and y coordinates.
func (p Point) IsOccupied(obstacles []Point) bool {
	for _, obs := range obstacles {
		if p == obs {
			return true
		}
	}
	return false
}

// NewPoint creates a new point with the given x and y coordinates.
func (p Point) Key() string {
	const delimiter = "-"
	return fmt.Sprint(p.X) + delimiter + fmt.Sprint(p.Y)
}

// NewPoint creates a new point with the given x and y coordinates.
func (p Point) IsEqual(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}
