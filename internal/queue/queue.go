package queue

import (
	"sync"
	"time"
)

// Default queue size to avoid unnecessary allocations
const defaultSize = 10

type Queue struct {
	mu    sync.Mutex
	items [][]byte

	// notify channel is used to notify the awaiting consumer that the queue is not empty
	notify chan struct{}
}

func NewQueue() *Queue {
	return &Queue{
		items:  make([][]byte, 0, defaultSize),
		notify: make(chan struct{}, 1),
	}
}

func (q *Queue) Push(item []byte) {
	q.mu.Lock()

	wasEmpty := len(q.items) == 0
	q.items = append(q.items, item)

	q.mu.Unlock()

	if wasEmpty {
		select {
		case q.notify <- struct{}{}:
		default:
		}
	}
}

func (q *Queue) Pop() []byte {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]

	copy(q.items, q.items[1:])
	q.items[len(q.items)-1] = nil
	q.items = q.items[:len(q.items)-1]

	return item
}

func (q *Queue) PopWait(timeout time.Duration) []byte {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		if item := q.Pop(); item != nil {
			return item
		}

		select {
		case <-q.notify:
		case <-timer.C:
			return nil
		}
	}
}
