package generic

import (
	"sync"
)

// Queue is a simple queue with unique items.
type Queue struct {
	queue    []Hashable
	keyCache map[interface{}]bool
	lock     sync.RWMutex
}

// NewQueue create a new Queue.
func NewQueue() Queue {
	return Queue{
		queue:    []Hashable{},
		keyCache: map[interface{}]bool{},
	}
}

// Push items into the queue.
// The queue contains only unique items.
func (q *Queue) Push(items ...Hashable) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, item := range items {
		if v, found := q.keyCache[item.Hash()]; found && v {
			continue
		}

		q.keyCache[item.Hash()] = true
		q.queue = append(q.queue, item)
	}
}

// Pull an item from the Queue.
func (q *Queue) Pull() Hashable {
	if len(q.queue) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()

		item := q.queue[0]
		q.queue = q.queue[1:]

		q.keyCache[item.Hash()] = false

		return item
	}

	return nil
}

// Size of the queue.
func (q *Queue) Size() int {
	return len(q.queue)
}

// Empty or not?
func (q *Queue) Empty() bool {
	return q.Size() == 0
}
