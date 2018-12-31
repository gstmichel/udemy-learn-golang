package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var cpt = 0

func main(){
	nb := 10
	wg.Add(nb)
	for i:= 0; i < nb; i++ {
		go routine()
	}
	wg.Wait()

	fmt.Println(cpt)
}

func routine(){
	x := cpt
	runtime.Gosched()
	x++
	cpt = x

	wg.Done()
}