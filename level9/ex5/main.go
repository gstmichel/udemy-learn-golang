package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var cpt int64 = 0

func main(){
	nb := 10
	wg.Add(nb)
	for i:= 0; i < nb; i++ {
		go routine()
	}
	wg.Wait()
}

func routine(){
	atomic.AddInt64(&cpt, 1)
	fmt.Println(atomic.LoadInt64(&cpt))
	wg.Done()
}