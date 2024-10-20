package linkedlistwithtail

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
		l.addFirstItem(item)
	} else {
		l.head = &linkedListItem{item: item, prev: l.head}
	}
}

func (l *linkedList) Append(item interface{}) {
	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.tail.prev = &linkedListItem{item: item, prev: nil}
		l.tail = l.tail.prev
	}
}

func (l *linkedList) RemoveHead() interface{} {
	if l.head == nil {
		return nil
	}

	removed := l.head.item
	l.head = l.head.prev
	return removed
}

func (l *linkedList) RemoveTail() interface{} {
	if l.head == nil {
		return nil
	}

	nextAfterTail := l.head

	for nextAfterTail.prev != l.tail {
		nextAfterTail = nextAfterTail.prev
	}

	nextAfterTail.prev = nil
	removed := l.tail.item
	l.tail = nextAfterTail
	return removed
}

func (l *linkedList) removeHeadOrTailNode(headOrTail *linkedListItem) {

}

func (l *linkedList) addFirstItem(item interface{}) {
	l.head = &linkedListItem{item: item}
	l.tail = l.head
}
