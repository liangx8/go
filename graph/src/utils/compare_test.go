package utils

import (
    "testing"
)

func TestCompareBuilder(t *testing.T) {
	var cb CompareBuilder
	cp := cb.Build()
	if _,err:=cp(nil,nil);err != nil {
		t.Fatal(err)
	}
}
func TestPrimitiveCompare(t *testing.T) {
	// string
	src,dst := "1","2"
	res,err :=PrimitiveCompare(src,dst)
	if err != nil {
		t.Error(err)
	}
	if res>0 {
		t.Errorf("input (%s,%s) expect result less than 0, result is %d",src,dst,res)
	}
	src,dst = "20","2"
	res,_ =PrimitiveCompare(src,dst)
	if res<0 {
		t.Errorf("input (%s,%s) expect result greater than 0, result is %d",src,dst,res)
	}
	src,dst = "2","20"
	res,_ =PrimitiveCompare(src,dst)
	if res>0 {
		t.Errorf("input (%s,%s) expect result less than 0, result is %d",src,dst,res)
	}
	src,dst = "0","20"
	res,_ =PrimitiveCompare(src,dst)
	if res>0 {
		t.Errorf("input (%s,%s) expect result less than 0, result is %d",src,dst,res)
	}
	src,dst = "100","20"
	res,_ =PrimitiveCompare(src,dst)
	if res>0 {
		t.Errorf("input (%s,%s) expect result less than 0, result is %d",src,dst,res)
	}
	src,dst = "100","100"
	res,_ =PrimitiveCompare(src,dst)
	if res != 0 {
		t.Errorf("input (%s,%s) expect result equal to 0, result is %d",src,dst,res)
	}
}


