package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var cpt = 0

func main() {
	nb := 10
	wg.Add(nb)
	for i := 0; i < nb; i++ {
		go routine()
	}
	wg.Wait()

	fmt.Println(cpt)
}

func routine() {
	mtx.Lock()
	x := cpt
	x++
	cpt = x
	mtx.Unlock()

	wg.Done()
}
