package collection

import (
    "testing"
)

func TestXYZ(t *testing.T) {
	ss := NewStack()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	it := ss.Iterator()
	v := 4
	for it.HasNext() {
		if v != it.Next() {
			t.Errorf("value is not correted order")
		}
		v --
	}

	t.Log(ss.Pop(),ss.Pop(),ss.Pop(),ss.Pop())
	if !ss.IsEmpty() {
		t.Errorf("stack should be empty")
	}
}


