package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	item []int
}

func (s *Stack) Push(value int) {
	s.item = append(s.item, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.item) == 0 {
		return 0, errors.New("Stack is empty")
	}
	index := len(s.item) - 1
	value := s.item[index]
	s.item = s.item[0:index]

	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.item) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.item[len(s.item)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.item) == 0
}

func main() {
	s := Stack{}
	s.Push(1)
	v, e := s.Pop()

	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Printf("poped value = %d \n", v)

	fmt.Printf("Is Empty = %t \n", s.IsEmpty())

}
