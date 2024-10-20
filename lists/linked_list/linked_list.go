package linkedlist

type linkedListItem struct {
	item interface{}
	prev *linkedListItem
}

type linkedList struct {
	head *linkedListItem
	tail *linkedListItem
}

func (l *linkedList) Prepend(item interface{}) {
	if l.head == nil {
		l.head = &linkedListItem{item: item}
		l.tail = l.head
	} else {
		l.head = &linkedListItem{item: item, prev: l.head}
	}
}
