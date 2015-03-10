package collection

import (
	"testing"
	"io"
)
type Int struct {
	V int
}
func TestOptr(t *testing.T) {
	ss := NewStack()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	it := ss.Iterator()
	v := 4
	for {
		var e interface{}
		if err:=it.Next(&e);err == io.EOF{ break }
		if v != e {
			t.Errorf("value is not correted order")
		}
		v --
	}

	ss.Pop(nil)
	ss.Pop(nil)
	ss.Pop(nil)
	ss.Pop(nil)
	if !ss.IsEmpty() {
		t.Errorf("stack should be empty")
	}
	if ss.Pop(nil)!= io.EOF {
		t.Errorf("empty stock return error")
	}
}

func TestType(t *testing.T){
	ss := NewStack()
	v := Int{1}
	ss.Push(&v)
	v.V=2
	var e *Int
	ss.Pop(&e)
	if v.V != e.V {
		t.Errorf("%v,%v, not expected",v,e)
	}
	var p interface{}
	ss.Push(v)
	ss.Pop(&p)
	t.Log(p)
	p=v
}
