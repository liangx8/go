package collection

import (
	"algorithm"
//	"fmt"
)

const iterate_eof="btree eof"
type node struct{
	e interface{}
	l,r *node
}
type bag struct{
	top *node
	cp algorithm.Compare
}
type iterForBtree struct {
	ch chan interface{}
	next interface{}
}
func (itb *iterForBtree)Next() interface{} {
	cur := itb.next
	itb.next = <-itb.ch
	return cur
}
func (itb *iterForBtree)HasNext() bool {
	return itb.next!=iterate_eof
}


func (b *bag)Iterator() Iterator{
	ch := make(chan interface{})
	go func(){
		travel(b.top,ch)
		ch <- iterate_eof
	}()
	return &iterForBtree{ch,<-ch}
}
func (b *bag)Add(e interface{}) error{
	if b.top == nil {
		b.top=&node{e,nil,nil}
		return nil
	}
	return add(b.top,e,b.cp)
}
func (b *bag)Remove(e interface{}){}
func (b *bag)IsEmpty()bool{return b.top == nil}

func NewBag(cp algorithm.Compare) Bag{return &bag{top:nil,cp:cp}}
func NewSimpleBag() Bag{return &bag{top:nil,cp:algorithm.PrimitiveCompare}}

func add(cur *node,e interface{},cp algorithm.Compare) error{
	res,err := cp(e,cur.e)
	if err != nil {return err}
	
	if res>0 {
		if cur.r == nil {
			cur.r = &node{e,nil,nil}
		} else {
			return add(cur.r,e,cp)
		}
	} else {
		if cur.l == nil {
			cur.l = &node{e,nil,nil}
		} else {
			return add(cur.l,e,cp)
		} 
	}
	return nil
}

func travel(top *node, ch chan interface{}){
	if top.l != nil {
		travel(top.l,ch)
	}
	ch <- top.e
	if top.r != nil {
		travel(top.r,ch)
	}
}