// Some sorting algorithm
package sortable

import (
	"errors"
)
// compare two object
// return ErrorNotComparable
type CompareType func(l,r interface{}) (int,error)

func IntCompare(l,r interface{}) (int,error){
	err := errors.New("Expecting int type")
	lv,ok := l.(int)
	if !ok { return 0,err}
	rv, ok := r.(int)
	if !ok { return 0,err}
	return lv-rv,nil
}

//var ErrorNotComparable=errors.New("Object is not compare")

