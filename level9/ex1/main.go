package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	wg.Add(2)
	go say1()
	go say2()
	wg.Wait()

	fmt.Println("Done")
}

func say1() {
	fmt.Println("Say 1")
	wg.Done()
}

func say2() {
	fmt.Println("Say 2")
	wg.Done()
}