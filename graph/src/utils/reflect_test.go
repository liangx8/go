package utils

import (
    "testing"
)

type S struct {
	V int
}

func TestCp(t *testing.T){
	i := 1
	pi := new(int)
	if err:=Cp(i,pi); err != nil {
		t.Fatal(err)
	}

	if i != *pi {
		t.Fatal("Cp() failed")
	}
	var s S
	ps := new(S)
	s.V=100
	if err:=Cp(s,ps); err != nil {
		t.Fatal(err)
	}
}


