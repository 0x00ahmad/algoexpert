package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	outSumLL := &LinkedList{}
	currSumNode := outSumLL
	carry := 0

	for linkedListOne != nil && linkedListTwo != nil {
		currSumNode.Next, carry = GetSummedNode(linkedListOne, linkedListTwo, carry)
		currSumNode = currSumNode.Next
		linkedListOne = linkedListOne.Next
		linkedListTwo = linkedListTwo.Next
	}

	if linkedListOne != nil {
		for linkedListOne != nil {
			currSumNode.Next = &LinkedList{linkedListOne.Value + carry, nil}
			currSumNode = currSumNode.Next
			linkedListOne = linkedListOne.Next
			carry = 0
		}
	}

	if linkedListTwo != nil {
		for linkedListTwo != nil {
			currSumNode.Next = &LinkedList{linkedListTwo.Value + carry, nil}
			currSumNode = currSumNode.Next
			linkedListTwo = linkedListTwo.Next
			carry = 0
		}
	}

	if carry != 0 {
		currSumNode.Next = &LinkedList{carry, nil}
	}

	return outSumLL.Next
}

func GetSummedNode(n1 *LinkedList, n2 *LinkedList, carry int) (*LinkedList, int) {
	s := n1.Value + n2.Value + carry
	modded_sum := s % 10
	c := 0
	if s > 9 {
		c = 1
	}
	return &LinkedList{modded_sum, nil}, c
}
