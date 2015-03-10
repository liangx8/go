package collection

import (
	"errors"
)

var ErrOutOfRange = errors.New("index is out of range!")

type Iterator interface{
	Next(interface{})bool
}
type Iteration interface{
	Iterator() Iterator
}
// FILO order
type Stack interface{
	Iteration
	Push(e interface{})
	Pop(interface{})error
	IsEmpty()bool
}
// Choas order
type Bag interface{
	Iteration
	Add(e interface{})
	Remove(e interface{})
	Find(e interface{})error
}
// Random access
type Vector interface{
	Iteration
	Add(e interface{})
	Set(idx int,e interface{}) error
	Get(idx int,e interface{}) error
	Remove(idx int,e interface{})error
	Size() int
}
func ForEach(c Iteration,cb func(e interface{})error)error{
	it := c.Iterator()
	var e interface{}
	for it.Next(&e){
		if err:=cb(e);err != nil {
			return err
		}
	}
	return nil
}
