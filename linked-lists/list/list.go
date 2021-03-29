package list

type List interface {
	InsertFront(interface{})
	//	InsertBack(int)
	RemoveFront()
	//	RemoveBack()
	Traverse()
	Size() int
	GetHead() interface{}
	GetTail() interface{}
}
