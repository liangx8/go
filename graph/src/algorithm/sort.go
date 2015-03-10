package algorithm

import (
	"collection"
	"utils"
	"io"
//	"fmt"
)

func SimpleQuickSort(es []interface{}) error{
	return QuickSort(es,utils.PrimitiveCompare)
}
func QuickSort(es []interface{},c utils.Compare) error{
//	ss := collection.NewStack()
	start,end := 0,len(es)
	return _quickSort(es,start,end,c)
}
func _quickSort(es []interface{},start,end int,c utils.Compare) error {
	var st,en int
	var dir_left bool
	var res int
	var err error


	if end - start < 2 {return nil}
	st,en =start,end-1
	dir_left =true
	for st<en{
		res,err =c(es[st],es[en])
		if err !=nil {return err}
		if dir_left {
			if res>0 {
				es[st],es[en]=es[en],es[st]
				st ++
				dir_left=false
			} else {
				en --
			}
		} else {
			if res>0 {
				es[st],es[en]=es[en],es[st]
				en--
				dir_left=true
			} else {
				st ++
			}
		}
	}
	err = _quickSort(es,start,st,c)
	if err != nil { return err }
	err = _quickSort(es,st+1,end,c)
	if err != nil { return err }
	return nil
}
func SimpleQsort(es []interface{}) error {
	return Qsort(es,utils.PrimitiveCompare)
}
/*
不用递归函数的递归写法
*/
func Qsort(es []interface{},c utils.Compare) error{
	start,end := 0, len(es)
	var st,en int
	var dir_left bool
	var res int
	var err error
	var sp *stack_pool
	ss := collection.NewStack()
start_:
	if end - start < 2 {
		goto middle_
	}
	st,en =start,end-1
	dir_left =true
	for st<en{
		res,err =c(es[st],es[en])
		if err !=nil {return err}
		if dir_left {
			if res>0 {
				es[st],es[en]=es[en],es[st]
				st ++
				dir_left=false
			} else {
				en --
			}
		} else {
			if res>0 {
				es[st],es[en]=es[en],es[st]
				en--
				dir_left=true
			} else {
				st ++
			}
		}
	}
	sp = &stack_pool{start,st,end,en,true}
	ss.Push(sp)
	end=st
	goto start_
middle_:
	//if ss.IsEmpty() {return nil}

	if ss.Pop(&sp) == io.EOF { return nil }
	start,st,end,en=sp.start,sp.st,sp.end,sp.en
	if !sp.left {goto middle_right}
//middle_left:
	sp.left=false
	ss.Push(sp)
	start=st+1

	goto start_
middle_right:
	goto middle_
}
type stack_pool struct {
	start,st,end,en int
	left bool
}
