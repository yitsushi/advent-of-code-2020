package day17

import (
	"sync"
)

type UpdateQueue struct {
	queue    []*Node
	keyCache map[string]bool
	lock     sync.RWMutex
}

func NewUpdateQueue() UpdateQueue {
	return UpdateQueue{
		queue:    []*Node{},
		keyCache: map[string]bool{},
	}
}

func (q *UpdateQueue) Push(nodes ...*Node) {
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

func (q *UpdateQueue) Pull() *Node {
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

func (q *UpdateQueue) Size() int {
	return len(q.queue)
}

func (q *UpdateQueue) Empty() bool {
	return len(q.queue) == 0
}
