/*
- 二叉树(binary tree)
- Go Channel

遍历树在协程中运行,结果输出到channel中,然后在主进程读取channel中的指,这是一个相对优雅的设计
*/
package collection

import (
	"utils"
	"fmt"
)
var NO_FOUND = fmt.Errorf("No found")
const iterate_eof="btree eof"
type node struct{
	e interface{}
	l,r *node
}
type bag struct{
	top *node
	cp utils.Compare
}
type iterForBtree struct {
	ch chan interface{}
	next interface{}
}
// it's unpredictable to call this method if HasNext() reutrn false
func (itb *iterForBtree)Next() interface{} {
	cur := itb.next.(*node).e
	itb.next = <-itb.ch
	return cur
}
func (itb *iterForBtree)HasNext() bool {
	return itb.next!=iterate_eof
}


func (b *bag)Find(e interface{}) error {
	if b.top == nil {
		return NO_FOUND
	}
	_,_,err := find_btree(nil,b.top,e,b.cp)
	return err
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
/*
删除节点到算法：
1 被删除的节点是左
  1 1 重新排序左节点的顺序
  1 2 排序后到树到根接到被删除的节点的位置
2 
3 
*/
func (b *bag)Remove(e interface{}){
	p,del,err := find_btree(nil,b.top,e,b.cp)
	if err != nil {return }
	if p== nil {
		b.top=del.r
	} else {
		p.r=del.r
	}
}
func (b *bag)IsEmpty()bool{return b.top == nil}

func NewBag(cp utils.Compare) Bag{return &bag{top:nil,cp:cp}}
func NewSimpleBag() Bag{return &bag{top:nil,cp:utils.PrimitiveCompare}}

func add(cur *node,e interface{},cp utils.Compare) error{
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

func travel(top *node, ch chan interface{}) {
	if top.l != nil {
		travel(top.l,ch)
	}
	ch <- top
	if top.r != nil {
		travel(top.r,ch)
	}
}
/*
* 增加父节点跟踪。提供目标节点的父节点
*/
func find_btree(parent,top *node,e interface{},cp utils.Compare)(*node,*node,error){
	res,err := cp(e,top.e)
	if err != nil {
		return nil,nil,err
	}
	if res < 0 {
		if top.l != nil {
			return find_btree(top,top.l,e,cp)
		}
		return nil,nil,NO_FOUND
	} else if res > 0 {
		if top.r != nil {
			return find_btree(top,top.r,e,cp)
		}
		return nil,nil,NO_FOUND
	}
	return parent,top,nil
}
func best_left(cur *node) *node {
	if cur.l != nil {
		return best_left(cur.l)
	}
	return cur
}
func CountDepth(b Bag) (int,error) {
	ba,ok := b.(*bag)
	if !ok {
		return 0,fmt.Errorf("it's not type collection.bag")
	}
	ch := make(chan int)
	go func() {
		countDepth(ba.top,ch)
		ch <- 99
	}()
	sum :=0
	for {
		r := <-ch
		if r == 99 {break}
		sum = sum + r
	}
	return sum,nil

}
func countDepth(top *node,ch chan int){
	if top.l != nil {
		ch <-1
		countDepth(top.l,ch)
	}
	if top.r != nil {
		ch <- -1
		countDepth(top.r,ch)
	}
}
