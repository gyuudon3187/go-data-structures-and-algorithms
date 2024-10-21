package linkedlist

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
		current := l.head
		for current.next != nil {
			current = current.next
		}

		current.next = &node{item: item}
	}

	l.len++
}

func (l *linkedList) RemoveHead() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	removed := l.removeHeadAndDecrementLength()
	return removed
}

func (l *linkedList) RemoveTail() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	beforeTail := l.head

	for beforeTail.next.next != nil {
		beforeTail = beforeTail.next
	}

	removed := beforeTail.next.item
	beforeTail.next = nil
	l.len--
	return removed
}

func (l *linkedList) RemoveAt(index int) (interface{}, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if index < 0 || index >= l.len {
		return nil, fmt.Errorf("Index out of bounds: index %d provided but list has length %d", index, l.len)
	}

	var removed interface{}

	if index == 0 {
		removed = l.removeHeadAndDecrementLength()
		return removed, nil
	}

	beforeRemovedNode := l.head
	for i := 0; i < index-1; i++ {
		beforeRemovedNode = beforeRemovedNode.next
	}

	removedNode := beforeRemovedNode.next
	removed = removedNode.item

	l.removeAndDecrementLength(beforeRemovedNode, removedNode)

	return removed, nil
}

func (l *linkedList) RemoveItem(item interface{}) (interface{}, int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	beforeRemovedNode := l.head
	if beforeRemovedNode == nil {
		return nil, -1, fmt.Errorf("Could not remove the following item because the list is empty: %v", item)
	}

	i := 0
	var removed interface{}

	if beforeRemovedNode.item == item {
		removed = l.removeHeadAndDecrementLength()
		return removed, i, nil
	}

	for beforeRemovedNode.next != nil {
		if beforeRemovedNode.next.item == item {
			removedNode := beforeRemovedNode.next
			removed = removedNode.item

			l.removeAndDecrementLength(beforeRemovedNode, removedNode)

			return removed, i, nil
		}

		beforeRemovedNode = beforeRemovedNode.next
		i++
	}

	return nil, -1, fmt.Errorf("No such item in the list: %v", item)
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
		} else {
			fmt.Printf(", %v", item)
		}

		fmt.Printf(")")
	})
}

func (l *linkedList) removeHeadAndDecrementLength() interface{} {
	removed := l.head.item
	l.head = l.head.next
	l.len--

	return removed
}

func (l *linkedList) removeAndDecrementLength(beforeNodeToBeRemoved, nodeToBeRemoved *node) {
	beforeNodeToBeRemoved.next = nodeToBeRemoved.next

	l.len--
}

func (l *linkedList) addFirstItem(item interface{}) {
	l.head = &node{item: item}
}
