package day17

import (
	"sync"
)

// NodeQueue is a simple queue with unique items.
type NodeQueue struct {
	queue    []*Node
	keyCache map[string]bool
	lock     sync.RWMutex
}

// NewUpdateQueue create a new UpdateQueue.
func NewUpdateQueue() NodeQueue {
	return NodeQueue{
		queue:    []*Node{},
		keyCache: map[string]bool{},
	}
}

// Push items into the queue.
// The queue contains only unique items.
func (q *NodeQueue) Push(nodes ...*Node) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, node := range nodes {
		if v, found := q.keyCache[node.Coordinate.Hash()]; found && v {
			continue
		}

		q.keyCache[node.Coordinate.Hash()] = true
		q.queue = append(q.queue, node)
	}
}

// Pull an item from the Queue.
func (q *NodeQueue) Pull() *Node {
	if len(q.queue) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()

		item := q.queue[0]
		q.queue = q.queue[1:]

		q.keyCache[item.Coordinate.Hash()] = false

		return item
	}

	return nil
}

// Size of the queue.
func (q *NodeQueue) Size() int {
	return len(q.queue)
}

// Empty or not?
func (q *NodeQueue) Empty() bool {
	return q.Size() == 0
}
