package sortable

import (
	"testing"
)

func Test_bubble(t *testing.T){
	ai := []interface{} {4,2,5,3,1,0}
	err := Bubble(ai,IntCompare)
	if err != nil {
		t.Fatal(err)
	}
	for i :=0 ;i<6 ;i++ {
		if ai[i] != i {
			t.Fatalf("compare is not corrected")
		}
	}
}
