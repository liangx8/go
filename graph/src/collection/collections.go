package collection

import (

)

type Iterator interface{
	HasNext() bool
	Next() interface{}
}
type Iteration interface{
	Iterator() Iterator
}
type Stack interface{
	Iteration
	Push(e interface{})
	Pop() interface{}
	IsEmpty()bool
}
// sequence access
type Bag interface{
	Iteration
	Add(e interface{}) error
	Remove(e interface{})
	Find(e interface{})error
}
// random access
type Vector interface{
	Iteration
	Add(e interface{})
	Set(idx int,e interface{})
	Get(idx int)interface{}
	Remove(idx int) interface{}
	Size() int
}
func ForEach(c Iteration,cb func(e interface{})){
	it := c.Iterator()
	for it.HasNext() {
		cb(it.Next())
	}
}
