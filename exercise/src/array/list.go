package array

import (
	"strings"
	"fmt"
)
type linkedElement struct{
	prev,next *linkedElement
	e interface{}
}
type List struct{
	head,tail *linkedElement
	size int
}
func (l *List)PushBack(e interface{}){
	le := new(linkedElement)
	le.next=l.tail
	le.prev=l.tail.prev
	l.tail.prev=le
	le.prev.next = le
	le.e=e
	l.size++
}
func initList(l *List){
	l.tail=new(linkedElement)
	l.tail.next=nil
	l.head=new(linkedElement)
	l.head.next=l.tail
	l.head.prev=nil
	l.tail.prev=l.head
}
func (l *List)PushFront(e interface{}){
	le := new(linkedElement)
	le.e=e
	l.size++
	le.prev=l.head
	le.next=l.head.next
	l.head.next=le
	le.next.prev=le
}
func (l *List)PopFront() interface{}{return nil}
func (l *List)PopBack() interface{}{return nil}
func (l *List)Size() int{return l.size}
func (l *List)String() string{
	var buf []string = make([]string,l.Size()+2)
	buf[0]="List["
	buf[l.Size()+1]="]"
	itr := l.head.next
	i := 0
	for itr != l.tail {
		buf[i+1]=fmt.Sprint(itr.e)
		i++
		itr = itr.next
	}
	return strings.Join(buf," ")
}
func (l *List)Clear(){
	l.head.next=l.tail
	l.tail.prev=l.head
	l.size=0
}
func NewList() List{
	l := List{nil,nil,0}
	initList(&l)
	return l
}
