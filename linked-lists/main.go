package main

import (
	"fmt"
	"github.com/nolancon/golang_algorithms/linked-lists/circular"
	"github.com/nolancon/golang_algorithms/linked-lists/double"
	"github.com/nolancon/golang_algorithms/linked-lists/list"
	"github.com/nolancon/golang_algorithms/linked-lists/single"
)

func main() {

	fmt.Printf("\n************\nSingly Linked Lists...\n************\n")
	var singleList list.List = &single.SinglyLinkedList{}

	singleList.InsertFront(5)
	singleList.Traverse()

	singleList.InsertFront(8)
	singleList.Traverse()

	singleList.InsertFront(7)
	singleList.Traverse()

	singleList.RemoveFront()
	singleList.Traverse()

	singleList.RemoveFront()
	singleList.Traverse()

	fmt.Println("\n************\nDoubly Linked Lists...\n************\n")
	var doubleList list.List = &double.DoublyLinkedList{}

	doubleList.InsertFront(200)
	doubleList.Traverse()

	doubleList.InsertFront(300)
	doubleList.Traverse()

	doubleList.InsertFront(100)
	doubleList.Traverse()

	doubleList.RemoveFront()
	doubleList.Traverse()

	doubleList.RemoveFront()
	doubleList.Traverse()

	fmt.Println("\n************\nCircularly Linked Lists...\n************\n")
	var circularList list.List = &circular.CircularlyLinkedList{}

	circularList.InsertFront("cat")
	circularList.Traverse()

	circularList.InsertFront("dog")
	circularList.Traverse()

	circularList.InsertFront("rabbit")
	circularList.Traverse()

	circularList.RemoveFront()
	circularList.Traverse()

	circularList.RemoveFront()
	circularList.Traverse()

}
