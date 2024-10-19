package queue

import "sync"

type queueItem struct {
	item interface{}
	prev *queueItem
}

type queue struct {
	first *queueItem
	last  *queueItem
	mu    sync.Mutex
}

func New() *queue {
	return new(queue)
}

func (q *queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.first == nil {
		q.first = &queueItem{item: item}
		q.last = q.first
	} else {
		q.last.prev = &queueItem{item: item}
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
