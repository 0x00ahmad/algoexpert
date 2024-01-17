package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	curr := ll.Head
	for curr != nil && curr.Value != value {
		curr = curr.Next
	}
	return curr != nil
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	curr := ll.Head
	currPos := 1
	for curr != nil && currPos != position {
		curr = curr.Next
		currPos += 1
	}
	if curr != nil {
		ll.InsertBefore(curr, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	curr := ll.Head
	for curr != nil {
		nodeToRemove := curr
		curr = curr.Next
		if nodeToRemove.Value == value {
			ll.Remove(nodeToRemove)
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	removeNodePointer(node)
}

func removeNodePointer(node *Node) {
	// re-points the previous and the next node's pointers to the node passed
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Next = nil
	node.Prev = nil
}
