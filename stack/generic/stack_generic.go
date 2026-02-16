package main

import (
	"errors"
	"fmt"
)

type GenericStack[T any] struct {
	item []T
}

func (s *GenericStack[T]) Push(i T) {
	s.item = append(s.item, i)
}

func (s *GenericStack[T]) Pop() (T, error) {
	var zero T
	if len(s.item) == 0 {
		return zero, errors.New("Empty stack")
	}
	index := len(s.item) - 1
	value := s.item[index]
	s.item = s.item[0:index]
	return value, nil
}

func (s *GenericStack[T]) Peek() (T, error) {
	var zero T
	if len(s.item) == 0 {
		return zero, errors.New("Empty stack")
	}
	index := len(s.item) - 1
	value := s.item[index]
	return value, nil
}

func (s *GenericStack[T]) IsEmpty() bool {
	return len(s.item) == 0
}

func main() {
	s := GenericStack[string]{}
	s.Push("one")
	fmt.Println(s.item)

}
