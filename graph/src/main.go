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
	for i:=0;i<13;i++ {
		b.Add(rand.Int31())
	}
	var 计数器 int
	collection.ForEach(b,func(e interface{}){
		fmt.Printf("%d %d\n",计数器,e)
		计数器++
	})
	
}

