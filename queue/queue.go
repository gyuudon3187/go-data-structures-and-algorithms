package queue

type queueItem struct {
	item interface{}
	prev *queueItem
}

type queue struct {
	first *queueItem
	last  *queueItem
}

func New() *queue {
	return new(queue)
}

func (q *queue) Enqueue(item interface{}) {
	if q.first == nil {
		q.first = &queueItem{item: item}
		q.last = q.first
	} else {
		q.last.prev = &queueItem{item: item}
		q.last = q.last.prev
	}
}

func (q *queue) Dequeue() interface{} {
	if q.first != nil {
		item := q.first.item
		q.first = q.first.prev
		return item
	}

	return nil
}
