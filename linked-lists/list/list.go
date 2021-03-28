package list

type List interface {
	InsertFront(int)
	//	InsertBack(int)
	RemoveFront()
	//	RemoveBack()
	Traverse()
	Size() int
}
