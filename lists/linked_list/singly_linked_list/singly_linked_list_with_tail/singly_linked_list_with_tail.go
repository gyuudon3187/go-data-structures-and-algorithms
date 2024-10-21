package linkedlistwithtail

import "fmt"

type node struct {
	item interface{}
	prev *node
}

type linkedList struct {
	head *node
	tail *node
	len  int
}

func (l *linkedList) Length() int {
	return l.len
}

func (l *linkedList) Prepend(item interface{}) {
	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.head = &node{item: item, prev: l.head}
	}

	l.len++
}

func (l *linkedList) Append(item interface{}) {
	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.tail.prev = &node{item: item, prev: nil}
		l.tail = l.tail.prev
	}

	l.len++
}

func (l *linkedList) RemoveHead() interface{} {
	if l.head == nil {
		return nil
	}

	removed := l.head.item
	l.head = l.head.prev
	l.len--
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
	l.len--
	return removed
}

func (l *linkedList) Find(item interface{}) *node {
	current := l.head

	for current != nil {
		if current.item == item {
			return current
		}

		current = current.prev
	}

	return nil
}

func (l *linkedList) IsEmpty() bool { return l.len == 0 }

func (l *linkedList) Iterate(action func(interface{})) {
	for node := l.head; node != nil; node = node.prev {
		action(node.item)
	}
}

func (l *linkedList) Print() {
	l.Iterate(func(item interface{}) {
		if l.head.item == item {
			fmt.Printf("[%v", item)
		} else if l.tail.item == item {
			fmt.Printf(", %v]", item)
		} else {
			fmt.Printf(", %v", item)
		}
	})
}

func (l *linkedList) addFirstItem(item interface{}) {
	l.head = &node{item: item}
	l.tail = l.head
}
