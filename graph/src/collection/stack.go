package collection

import (

)
type stack_node struct {
	e interface{}
	next *stack_node
}
type stack_iterator struct{
	itr *stack_node
}
func (si *stack_iterator)HasNext() bool{
	return si.itr != nil
}
func (si *stack_iterator)Next() interface{}{
	e := si.itr.e
	si.itr=si.itr.next
	return e
}
type stack_ struct{
	top *stack_node
}
func (s *stack_)Iterator() Iterator{
	return &stack_iterator{s.top}
}
func (s *stack_)Push(e interface{}){
	sn := &stack_node{e,s.top}
	s.top=sn
}
func (s *stack_)Pop() interface{}{
	e := s.top.e
	s.top=s.top.next
	return e
}
func (s stack_)IsEmpty()bool{
	return s.top==nil
}


func NewStack() Stack{
	return &stack_{nil}
}
