package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	head := linkedList
	current := linkedList

	for current != nil {
		if current.Next != nil && current.Value == current.Next.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

