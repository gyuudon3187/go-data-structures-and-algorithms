package linkedlistwithtail

import (
	"fmt"
	"sync"
)

type node struct {
	item interface{}
	next *node
}

type linkedList struct {
	head *node
	tail *node
	len  int
	mu   sync.Mutex
}

func New() *linkedList {
	return new(linkedList)
}

func (l *linkedList) Length() int {
	return l.len
}

func (l *linkedList) Prepend(item interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.head = &node{item: item, next: l.head}
	}

	l.len++
}

func (l *linkedList) Append(item interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.tail.next = &node{item: item}
		l.tail = l.tail.next
	}

	l.len++
}

func (l *linkedList) RemoveHead() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	removed := l.head.item
	l.head = l.head.next
	l.len--
	return removed
}

func (l *linkedList) RemoveTail() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	beforeTail := l.head

	for beforeTail.next != l.tail {
		beforeTail = beforeTail.next
	}

	beforeTail.next = nil
	removed := l.tail.item
	l.tail = beforeTail
	l.len--
	return removed
}

func (l *linkedList) Find(item interface{}) *node {
	current := l.head

	for current != nil {
		if current.item == item {
			return current
		}

		current = current.next
	}

	return nil
}

func (l *linkedList) IsEmpty() bool { return l.len == 0 }

func (l *linkedList) Iterate(action func(interface{})) {
	for node := l.head; node != nil; node = node.next {
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
