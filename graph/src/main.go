package main

import (
	"time"
	"math/rand"
	"fmt"

	"collection"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	b := collection.NewSimpleBag()
	for i:=0;i<20;i++ {
		var num int32
		for {
			num = rand.Int31()%20
			if b.Find(num) == collection.NO_FOUND {break}
		}
		b.Add(num)
	}
	fmt.Println(collection.CountDepth(b))
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

