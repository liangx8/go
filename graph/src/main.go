package main

import (
	"time"
	"math/rand"
	"fmt"


	"algorithm"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ia := make([]interface{},amount)
	for i:=0;i<amount;i++ {
		var num int32
		num = rand.Int31()%99
		ia[i]=num
	}
	fmt.Println(ia)
	err := algorithm.SimpleQsort(ia)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ia)
	fmt.Println("========================================")
/*
	var 计数器 int
	collection.ForEach(b,func(e interface{}){
		fmt.Printf("%d %d\n",计数器,e)
		计数器++
	})
	b.Remove(int32(9))
	fmt.Println("----------------------------------------")
	collection.ForEach(b,func(e interface{}){
		fmt.Printf("%d %d\n",计数器,e)
		计数器++
	})
*/

}
const (
	amount = 10
)
