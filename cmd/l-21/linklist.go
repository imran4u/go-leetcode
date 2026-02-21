package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkList struct {
	Head *ListNode
}

func (l *LinkList) Insert(val int) {
	node := &ListNode{Val: val}
	if l.Head == nil {
		l.Head = node
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkList) Println() {
	var sb strings.Builder

	if l.Head == nil {
		sb.WriteString("[]")
	} else {
		current := l.Head
		for current != nil {
			sb.WriteString(strconv.Itoa(current.Val))
			sb.WriteString("->")
			current = current.Next
		}
		sb.WriteString("nil")
	}

	fmt.Println(sb.String())

}

func (l *LinkList) MergeSortedList(list2 *LinkList) LinkList {
	var result LinkList

	dummy := &ListNode{}
	tail := dummy

	p1 := l.Head
	p2 := list2.Head

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}

	// Attach remaining nodes
	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}

	result.Head = dummy.Next
	return result
}
