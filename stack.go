package main

//
// The stack provides a way of switching views/levels/etc
// You push on to the stack the state you want to execute
// and once done you can remove it. Useful for things like
// pause screens or popup menus.
//

import (
	"errors"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []interface{}
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]interface{}, 0)}
}

func (s stack) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}

	return s.s[s.Len()-1]
}

func (s stack) Len() int {
	return len(s.s)
}

func (s stack) Empty() bool {
	return s.Len() == 0
}

func (s *stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
