package internal

import "log"

// printhPath prints the path from the start to the end point based on the visited points.
func printhPath(start Point, current queueItem, height int, width int) {
	var path string
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if current.IsVisited(Point{x, y}) {
				if start.X == x && start.Y == y {
					path += "S"
				} else {
					path += "*"
				}
			} else {
				path += "."
			}
		}
		path += "\n"
	}

	log.Printf("We reached the end point with velocity (%d, %d)\n", current.velocity.dx, current.velocity.dy)
	log.Printf("Hops: %d\n", current.hops)
	log.Printf("Shortes path: \n%s", path)
}

func PrintGrid(h, w int, start, finish Point, obstacles []Point) {
	var grid string

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := Point{x, y}
			switch {
			case p.IsEqual(start):
				grid += "S"
			case p.IsEqual(finish):
				grid += "F"
			case p.IsOccupied(obstacles):
				grid += "X"
			default:
				grid += "."
			}
		}
		grid += "\n"
	}

	log.Printf("Grid: \n%s", grid)
}
