package collection

import (
	"io"


	"utils"
)
type stack_node struct {
	e interface{}
	next *stack_node
}
type stack_iterator struct{
	itr *stack_node
}
// implement Iteratror
func (si *stack_iterator)Next(holder interface{})bool{
	if si.itr == nil {
		return false
	}
	e := si.itr.e
	si.itr=si.itr.next
	if err:=utils.Cp(e,holder) ; err != nil {
		panic(err)
	}
	return true
}
type stack_ struct{
	top *stack_node
}
// implement Iteration
func (s *stack_)Iterator() Iterator{
	return &stack_iterator{s.top}
}
func (s *stack_)Push(e interface{}) {
	sn := &stack_node{e,s.top}
	s.top=sn
}
func (s *stack_)Pop(holder interface{}) error{
	if s.top == nil {
		return io.EOF
	}
	e := s.top.e
	s.top=s.top.next
	if holder != nil{
		return utils.Cp(e,holder)
	}
	return nil
}
func (s stack_)IsEmpty()bool{
	return s.top==nil
}


func NewStack() Stack{
	return &stack_{nil}
}
