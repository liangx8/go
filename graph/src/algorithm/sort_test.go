package algorithm

import (
	"testing"
	"sort"

	"utils"
)

type sortable []interface{}


func (s sortable)Len() int {
	return len(s)
}
func (s sortable)Less(i,j int) bool {
	res,_ := utils.PrimitiveCompare(s[i],s[j])
	return res < 0
}
func (s *sortable)Swap(i,j int) {
	(*s)[i],(*s)[j]=(*s)[j],(*s)[i]
}

func TestQsort(t *testing.T) {
	ia := sortable{1,3,44,1,33,23,432,3,443,2,365,84,33,5,334,123,21,23}
	SimpleQsort(ia)
	if !sort.IsSorted(&ia) {
		t.Errorf("sort2 fail")
	}
}
func TestQuickSort(t *testing.T) {
	ia := sortable{1,3,44,1,33,23,432,3,443,2,365,84,33,5,334,123,21,23}
	SimpleQuickSort(ia)
	if !sort.IsSorted(&ia) {
		t.Errorf("sort1 fail")
	}
}

