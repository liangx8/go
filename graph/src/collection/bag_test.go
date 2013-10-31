package collection

import (
    "testing"
    "fmt"
)
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
	return nil
}
