package ds

import "errors"

type Stack[T any] []T

// Errors
var ErrEmptyStack = errors.New("stack is empty")

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var blank T
		return blank, ErrEmptyStack
	}
	removed := (*s)[0]
	*s = (*s)[1:]
	return removed, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(*s) == 0 {
		var blank T
		return blank, ErrEmptyStack
	}
	return (*s)[0], nil
}

func (s *Stack[T]) Push(item T) {
	*s = append([]T{item}, *s...)
}
