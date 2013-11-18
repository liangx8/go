package collection

import (
	"testing"
	"fmt"
	"utils"
)
func Test_btree(t *testing.T){
	num :=[]int{8,15,7,10,18,16,9}
	top := &node{num[0],nil,nil}
	for i:=1;i<len(num);i++{
		add(top,num[i],utils.PrimitiveCompare)
	}
	if top.l == nil || top.r == nil || top.r.l == nil || top.r.r == nil || top.r.l.l == nil || top.r.r.l == nil {
		t.Errorf("nodes are not in order")
		t.FailNow()
	}
	if top.r.e != num[1] || top.l.e != num[2] || top.r.l.e != num[3] || top.r.r.e != num[4] || top.r.r.l.e != num[5] || top.r.l.l.e != num[6] {
		t.Errorf("values are not correct")
	}
	ch := make(chan interface{})
	go travel(top,ch)
	if top.l != <-ch || top != <-ch || top.r.l.l != <-ch || top.r.l != <-ch || top.r != <-ch || top.r.r.l != <-ch || top.r.r != <-ch{
		t.Errorf("travel error")
	}
	p,c,_ := find_btree(nil,top,15,utils.PrimitiveCompare)
	if p != top || c != top.r {
		t.Errorf("level1 find error")
	}
	p,c,_ = find_btree(nil,top,9,utils.PrimitiveCompare)
	if p != top.r.l || c != top.r.l.l {
		t.Errorf("level2 find error")
	}
	t.Log("Test remove element")

}
func TestBag(t *testing.T) {
	err := testBag(NewSimpleBag())
	if err != nil {
		t.Error(err)
	}
}

func testBag(b Bag) error{

	s1 := "1first"
	s2 := "2second"
	s3 := "3third"
	s4 := "4forth"
	b.Add(s1)
	err := b.Add(1)
	if err == nil {
		return fmt.Errorf("input int(1), the first element is string,expect a error occurred!")
	}
	it := b.Iterator()
	if it == nil {
		return fmt.Errorf("expect .Iterator() return a object but nil")
	}
	b.Add(s3)
	b.Add(s2)
	b.Add(s4)
	b.Add(s4)
	it = b.Iterator()
	bRes:=it.HasNext()
	if !bRes{
		return fmt.Errorf("expect true but %v",bRes)
	}
	sRes := it.Next()
	if s1 != sRes {
		return fmt.Errorf("expect 1st element is %s but %s present",s1,sRes)
	}
	bRes =it.HasNext()
	if !bRes{
		return fmt.Errorf("expect true but %v",bRes)
	}
	sRes = it.Next()
	if s2 != sRes {
		return fmt.Errorf("expect 2nd element is %s but %s present",s2,sRes)
	}
	bRes =it.HasNext()
	if !bRes{
		return fmt.Errorf("expect true but %v",bRes)
	}
	sRes = it.Next()
	if s3 != sRes {
		return fmt.Errorf("expect 3th element is %s but %s present",s3,sRes)
	}
	bRes =it.HasNext()
	if !bRes{
		return fmt.Errorf("expect true but %v",bRes)
	}
	sRes = it.Next()
	if s4 != sRes {
		return fmt.Errorf("expect 4th element is %s but %s present",s4,sRes)
	}
	bRes =it.HasNext()
	if !bRes{
		return fmt.Errorf("expect true but %v",bRes)
	}
	sRes = it.Next()
	if s4 != sRes {
		return fmt.Errorf("expect 4th element is %s but %s present",s4,sRes)
	}
	bRes = it.HasNext()
	if bRes {
		return fmt.Errorf("expect false but %v",bRes)
	}
	err = b.Find(s1)
	if err == NO_FOUND {
		return fmt.Errorf("expect %s in %v,but no found",s1,b)
	}
	err = b.Find(100)
	if err == nil || err == NO_FOUND {
		return fmt.Errorf("try find a deference type but result is no error")
	}
	b.Remove(s1)
	err = b.Find(s1)
	if err == nil {
		return fmt.Errorf("%s was delete, but still found",s1)
	}

	return nil
}
