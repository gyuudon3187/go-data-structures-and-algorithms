package queue

import "sync"

type node struct {
	item interface{}
	prev *node
}

type queue struct {
	first *node
	last  *node
	mu    sync.Mutex
}

func New() *queue {
	return new(queue)
}

func (q *queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.first == nil {
		q.first = &node{item: item}
		q.last = q.first
	} else {
		q.last.prev = &node{item: item}
		q.last = q.last.prev
	}
}

func (q *queue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.first != nil {
		item := q.first.item
		q.first = q.first.prev
		return item
	}

	return nil
}
