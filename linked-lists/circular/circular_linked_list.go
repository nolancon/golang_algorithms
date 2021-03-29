package circular

import (
	"fmt"
)

type CircularlyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	data     interface{}
	next     *Node
	previous *Node
}

func (cl *CircularlyLinkedList) InsertFront(nodeData interface{}) {
	fmt.Printf("inserting %v to front of list...\n", nodeData)
	n := &Node{
		data: nodeData,
	}
	if cl.head == nil {
		n.previous = n
		n.next = n
		cl.head = n
		cl.tail = n

	} else {
		n.next = cl.head
		n.previous = cl.tail
		cl.head = n
		cl.head.next.previous = n
		cl.tail.next = n
	}

	cl.length++
}

func (cl *CircularlyLinkedList) RemoveFront() {
	fmt.Printf("removing head from list...\n")
	if cl.head == nil {
		return
	}
	if cl.head.next == nil {
		cl.head = nil
		cl.length--
		return
	}
	cl.head = cl.head.next
	cl.tail.next = cl.head
	cl.head.previous = cl.tail
	cl.length--
	return
}

func (cl *CircularlyLinkedList) Traverse() {
	if cl.head == nil {
		fmt.Printf("no head on list\n")
		return
	}
	fmt.Printf("traversing list...\n")
	currentNode := cl.head
	for {
		fmt.Printf("node...\n")
		if currentNode.previous != nil {
			fmt.Printf("previous node data: %v\n", currentNode.previous.data)
		}
		fmt.Printf("current node data: %v\n", currentNode.data)
		if currentNode.next != nil {
			fmt.Printf("next node data %v\n", currentNode.next.data)
		}
		if currentNode == cl.tail {
			break
		}
		currentNode = currentNode.next
	}
}

func (cl *CircularlyLinkedList) Size() int {
	return cl.length
}

func (cl *CircularlyLinkedList) GetHead() interface{} {
	return cl.head.data
}

func (cl *CircularlyLinkedList) GetTail() interface{} {
	return cl.tail.data
}
