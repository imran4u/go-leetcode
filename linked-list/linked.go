package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkedList struct {
	head *Node
}
type Node struct {
	data int
	next *Node
}

func (l *LinkedList) PostInsert(value int) {
	node := &Node{data: value}
	if l.head == nil {
		l.head = node
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = node

}

func (l *LinkedList) PreInsert(value int) {
	node := Node{value, nil}
	temp := l.head
	l.head = &node
	l.head.next = temp

}

func (l *LinkedList) deleteByValue(value int) {
	if l.head == nil {
		fmt.Println("Empty list nothing to delte")
		return
	}

	//delete head
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	// remove after
	pre := l.head
	current := l.head.next

	for current != nil {
		if current.data == value {
			pre.next = current.next
			return
		}
		current = current.next
		pre = pre.next
	}

	fmt.Println("value not found")
}

func (l LinkedList) Print() {

	var sb strings.Builder
	current := l.head
	for current != nil {
		sb.WriteString(strconv.Itoa(current.data))
		sb.WriteString("->")
		current = current.next
	}
	sbStr := sb.String() // convert to string
	if len(sbStr) > 2 {
		sbStr = sbStr[:len(sbStr)-2] // remove last -> I am sure it present
	}

	fmt.Println(sbStr)

}
func main() {
	list := LinkedList{}
	list.Print()
	fmt.Println(list) // nil
	//Preinsert
	list.PreInsert(1)
	list.PreInsert(2)
	list.PreInsert(3)
	fmt.Println(list) // memory address
	list.Print()

	// post insert
	list2 := LinkedList{}
	list2.PostInsert(1)
	list2.PostInsert(2)
	list2.PostInsert(3)
	list2.PostInsert(4)
	list2.Print()
	fmt.Println("After delete")
	list2.deleteByValue(4)
	list2.Print()
}
