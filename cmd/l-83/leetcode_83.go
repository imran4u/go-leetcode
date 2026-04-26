package main

import "fmt"

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Input: head = [1,1,2]
Output: [1,2]
Input: head = [1,1,2,3,3]
Output: [1,2,3]

*/

// if duplicate move current next otherwise move current
func deleteDuplicates(head *ListNode) *ListNode {

	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func main() {

	l := LinkedList{}
	l.insert(1)
	l.insert(1)
	l.insert(1)
	l.print()
	// delete duplicate
	fmt.Println("Delete duplicate")
	deleteDuplicates(l.head)
	l.print()

}
