package internal

import (
	"log"
)

// isValidGrid checks if the grid dimensions, start and finish points, and obstacles are valid.
func isValidGrid(start, finish Point, obstacles []Point, width, height int) bool {
	const maxGridSize = 30

	// Grid dimensions must be positive
	if width < 0 || height < 0 {
		return false
	}

	// Grid dimensions must be less than or equal to 30
	if width > maxGridSize || height > maxGridSize {
		return false
	}

	// Start and finish points must not be occupied by obstacles
	if start.IsOccupied(obstacles) || finish.IsOccupied(obstacles) {
		return false
	}

	// Start and finish points must be within the grid dimensions
	if start.X < 0 || start.X >= width || start.Y < 0 || start.Y >= height {
		return false
	}

	return true
}

// BFS is a function that returns the shortest path from start to finish
// given a grid with obstacles. It uses a breadth-first search algorithm
// to find the shortest path and returns the number of hops from start to
// finish. If there is no solution, it returns -1.
func BFS(start, finish Point, obstacles []Point, width, height int) int {
	if start == finish {
		return 0
	}

	if !isValidGrid(start, finish, obstacles, width, height) {
		return -1
	}

	// Create a queue with the start position and velocity (0, 0)
	queue := newQueue(start)
	visited := newVisitedGen()

	for !queue.isEmpty() {
		current := queue.dequeue()
		log.Printf("Elements in queue: %v", queue.Count())

		if current.position.IsEqual(finish) { // Hooray! We reached the end point
			printhPath(start, current, height, width)
			return current.hops // Return the number of hops
		}

		visited.add(current.hops, current.position)

		// Generate velocities within the restricted range
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				newVelocity := Velocity{current.velocity.dx + dx, current.velocity.dy + dy}
				newPosition := Point{current.position.X + newVelocity.dx, current.position.Y + newVelocity.dy}

				// Check if the new position is within the grid, not occupied by an obstacle,
				// not visited and the new velocity is within the restricted range.
				// If all conditions are met, enqueue the new position and velocity.
				if newVelocity.IsWithinRange() &&
					newPosition.IsWithinRange(width, height) &&
					!newPosition.IsOccupied(obstacles) &&
					!current.IsVisited(newPosition) &&
					!visited.isVisited(current.hops, newPosition) {

					queue.enqueue(queueItem{newPosition, newVelocity, current.hops + 1, current.visited})
				}
			}
		}
	}

	return -1 // No solution
}

type visitedGen struct {
	visited map[int]map[string]struct{}
}

func newVisitedGen() *visitedGen {
	return &visitedGen{visited: map[int]map[string]struct{}{}}
}

func (vg *visitedGen) add(gen int, p Point) {
	if vg.visited[gen] == nil {
		vg.visited[gen] = map[string]struct{}{}
	}

	vg.visited[gen][p.Key()] = struct{}{}
}

func (vg *visitedGen) isVisited(gen int, p Point) bool {
	newGen := gen
	if newGen < 0 {
		return false
	}

	if vg.visited[newGen] == nil {
		return false
	}

	for g := newGen; g >= 0; g-- {
		if _, ok := vg.visited[g][p.Key()]; ok {
			return true
		}
	}

	return false
}
