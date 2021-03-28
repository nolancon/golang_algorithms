package double

import (
	"fmt"
)

type DoublyLinkedList struct {
	head   *Node
	length int
}

type Node struct {
	data     int
	next     *Node
	previous *Node
}

func (dl *DoublyLinkedList) InsertFront(nodeData int) {
	fmt.Printf("inserting %v to front of list...\n", nodeData)
	n := &Node{
		data: nodeData,
	}
	if dl.head == nil {
		dl.head = n
	} else {
		n.next = dl.head
		dl.head = n
		dl.head.next.previous = n
	}

	dl.length++
}

func (dl *DoublyLinkedList) RemoveFront() {
	fmt.Printf("removing head from list...\n")
	if dl.head == nil {
		return
	}
	if dl.head.next == nil {
		dl.head = nil
		dl.length--
		return
	}
	dl.head = dl.head.next
	dl.head.previous = nil
	dl.length--
	return
}

func (dl *DoublyLinkedList) Traverse() {
	if dl.head == nil {
		fmt.Printf("no head on list\n")
		return
	}
	fmt.Printf("traversing list...\n")
	currentNode := dl.head
	for currentNode != nil {
		fmt.Printf("node...\n")
		if currentNode.previous != nil {
			fmt.Printf("previous node data: %v\n", currentNode.previous.data)
		}
		fmt.Printf("current node data: %v\n", currentNode.data)
		if currentNode.next != nil {
			fmt.Printf("next node data %v\n", currentNode.next.data)
		}
		currentNode = currentNode.next
	}
}
func (dl *DoublyLinkedList) Size() int {
	return dl.length
}
