package doublylinkedlist

import (
	"fmt"
	"sync"
)

type node struct {
	item interface{}
	next *node
	prev *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	len  int
	mu   sync.Mutex
}

func New() *doublyLinkedList {
	return new(doublyLinkedList)
}

func (l *doublyLinkedList) Length() int {
	return l.len
}

func (l *doublyLinkedList) Prepend(item interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.head = &node{item: item, next: l.head}
		l.head.next.prev = l.head
	}

	l.len++
}

func (l *doublyLinkedList) Append(item interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		l.addFirstItem(item)
	} else {
		l.tail.next = &node{item: item, prev: l.tail}
		l.tail = l.tail.next
	}

	l.len++
}

func (l *doublyLinkedList) RemoveHead() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	removed := l.removeHeadAndDecrementLength()
	return removed
}

func (l *doublyLinkedList) RemoveTail() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		return nil
	}

	removed := l.tail.item
	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
	return removed
}

func (l *doublyLinkedList) RemoveAt(index int) (interface{}, error) {
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

	l.setTailIfNewTailElseRemoveAndDecrement(beforeRemovedNode, removedNode)

	return removed, nil
}

func (l *doublyLinkedList) RemoveItem(item interface{}) (interface{}, int, error) {
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

			l.setTailIfNewTailElseRemoveAndDecrement(beforeRemovedNode, removedNode)

			return removed, i, nil
		}

		beforeRemovedNode = beforeRemovedNode.next
		i++
	}

	return nil, -1, fmt.Errorf("No such item in the list: %v", item)
}

func (l *doublyLinkedList) Find(item interface{}) *node {
	current := l.head

	for current != nil {
		if current.item == item {
			return current
		}

		current = current.next
	}

	return nil
}

func (l *doublyLinkedList) IsEmpty() bool { return l.len == 0 }

func (l *doublyLinkedList) Iterate(action func(interface{})) {
	for node := l.head; node != nil; node = node.next {
		action(node.item)
	}
}

func (l *doublyLinkedList) Print() {
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

func (l *doublyLinkedList) removeHeadAndDecrementLength() interface{} {
	removed := l.head.item
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.len--

	return removed
}

func (l *doublyLinkedList) setTailIfNewTailElseRemoveAndDecrement(beforeNodeToBeRemoved, nodeToBeRemoved *node) {
	if nodeToBeRemoved.next == nil {
		l.tail = beforeNodeToBeRemoved
		beforeNodeToBeRemoved.next = nil
	} else {
		beforeNodeToBeRemoved.next = nodeToBeRemoved.next
	}

	l.len--
}

func (l *doublyLinkedList) addFirstItem(item interface{}) {
	l.head = &node{item: item}
	l.tail = l.head
}
