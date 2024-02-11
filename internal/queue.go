package internal

import (
	"sync"

	"golang.org/x/exp/maps"
)

// queueItem represents an item in the queue.
type queueItem struct {
	position Point               // Current position
	velocity Velocity            // Current velocity
	hops     int                 // Number of hops
	visited  map[string]struct{} // Visited points
}

// IsVisited checks if the point has been visited.
func (qi *queueItem) IsVisited(p Point) bool {
	_, ok := qi.visited[p.Key()]
	return ok
}

// queue represents a queue of items.
type queue struct {
	sync.RWMutex
	items []queueItem
}

// enqueue adds an item to the queue.
func (q *queue) enqueue(item queueItem) {
	q.Lock()
	defer q.Unlock()

	visited := make(map[string]struct{})
	maps.Copy(visited, item.visited)
	visited[item.position.Key()] = struct{}{}
	item.visited = visited

	q.items = append(q.items, item)
}

// dequeue removes an item from the queue.
func (q *queue) dequeue() queueItem {
	q.Lock()
	defer q.Unlock()

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// isEmpty checks if the queue is empty.
func (q *queue) isEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.items) == 0
}

// newQueue creates a new queue with the given start position.
func newQueue(start Point) *queue {
	return &queue{items: []queueItem{{position: start, velocity: Velocity{0, 0}, hops: 0, visited: map[string]struct{}{start.Key(): {}}}}}
}

// Count returns the number of items in the queue.
func (q *queue) Count() int {
	q.RLock()
	defer q.RUnlock()
	return len(q.items)
}
