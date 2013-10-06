package main
import (
	"array"
	"fmt"
	"math"
)
func main(){
	var c array.Collection
	var s array.Sortable
	a := array.NewVector(0,10)
	c = &a
	s = &a
	c.PushBack("sentence")
	c.PushBack("文字")
	c.PushBack(math.Pi)
	c.PushFront("first")
	fmt.Printf("%d,\t%v\n",c.Size(),c)
	fmt.Printf("%v,\t%v\n",s.ElementOf(0),s)
	fmt.Println(c.PopFront(),c)
	fmt.Println(c.PopBack(),c)
	fmt.Println("=============================")

	l := array.NewList()
	c = &l
	c.PushBack("ListLast")
	c.PushFront(math.Pi)
	c.PushFront("ListFront")
	fmt.Printf("%d,%v\n",c.Size(),c)
	c.Clear()
	fmt.Printf("%d,%v\n",c.Size(),c)

	t := []int{1,2,3,4,5,6,7}
	fmt.Println(t[:len(t)-1])
}
