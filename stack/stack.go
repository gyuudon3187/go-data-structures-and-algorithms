package stack

import "sync"

type stackItem struct {
	item interface{}
	next *stackItem
}

type stack struct {
	sp *stackItem
	mu sync.Mutex
}

func New() *stack {
	return new(stack)
}

func (s *stack) Push(item interface{}) {
	s.sp = &stackItem{item, s.sp}
}

func (s *stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sp.next != nil {
		item := s.sp.item
		s.sp = s.sp.next
		return item
	}

	return nil
}

func (s *stack) Peek() interface{} {
	if s.sp != nil {
		return s.sp.item
	}

	return nil
}