package main

func main() {
	list1 := LinkList{}
	list1.Insert(1)
	list1.Insert(2)
	list1.Insert(4)
	list1.Println()

	list2 := LinkList{}
	list2.Insert(1)
	list2.Insert(3)
	list2.Insert(4)
	list2.Println()

	// merge
	result := list1.MergeSortedList(&list2)
	result.Println()
}
