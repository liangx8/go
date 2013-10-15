package array

type Collection interface{
	PushBack(e interface{})
	PushFront(e interface{})
	PopFront() interface{}
	PopBack() interface{}
	Clear()
	Iterator() Iterator
	Size() int
}
type Comparator func(l,r interface{}) int
type Sortable interface{
	Sort(c Comparator)
	ElementOf(index int) interface{}
}
type Iterator interface{
	HasNext() bool
	Next() interface{}
}
