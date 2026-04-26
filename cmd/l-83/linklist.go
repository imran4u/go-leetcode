package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) insert(val int) {
	node := ListNode{Val: val}
	if l.head == nil {
		l.head = &node
		return
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &node
}

func (l *LinkedList) print() {
	if l.head == nil {
		fmt.Println("Empty list")
	}
	current := l.head
	var sb strings.Builder
	for current != nil {
		sb.WriteString(fmt.Sprintf("%d->", current.Val))
		current = current.Next
	}
	sb.WriteString("nil")
	fmt.Println(sb.String())
}
