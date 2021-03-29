package single

import (
	"fmt"
)

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	data interface{}
	next *Node
}

func (sl *SinglyLinkedList) InsertFront(nodeData interface{}) {
	fmt.Printf("inserting %v to front of list...\n", nodeData)
	n := &Node{
		data: nodeData,
	}
	if sl.head == nil {
		sl.head = n
		sl.tail = n
	} else {
		n.next = sl.head
		sl.head = n
	}
	sl.length++
}

func (sl *SinglyLinkedList) RemoveFront() {
	fmt.Printf("removing head from list...\n")
	if sl.head == nil {
		return
	}
	if sl.head.next == nil {
		sl.head = nil
		sl.tail = nil
		sl.length--
		return
	}
	sl.head = sl.head.next
	sl.length--
	return
}

func (sl *SinglyLinkedList) Traverse() {
	if sl.head == nil {
		fmt.Printf("no head on list\n")
		return
	}
	fmt.Printf("traversing list...\n")
	currentNode := sl.head
	for currentNode != nil {
		fmt.Printf("node data: %v\n", currentNode.data)
		currentNode = currentNode.next
	}
}

func (sl *SinglyLinkedList) Size() int {
	return sl.length
}

func (sl *SinglyLinkedList) GetHead() interface{} {
	return sl.head.data
}

func (sl *SinglyLinkedList) GetTail() interface{} {
	return sl.tail.data
}
