package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	javaScriptPointer := head
	blazinglyFastPointer := head
	for i := 0; i < k; i++ {
		blazinglyFastPointer = blazinglyFastPointer.Next
	}
	if blazinglyFastPointer == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}
	for blazinglyFastPointer != nil && blazinglyFastPointer.Next != nil {
		javaScriptPointer = javaScriptPointer.Next
		blazinglyFastPointer = blazinglyFastPointer.Next
	}
	javaScriptPointer.Next = javaScriptPointer.Next.Next
}
