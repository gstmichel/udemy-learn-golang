package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- i*100 + j
			}
			wg.Done()
		}(i)
	}

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	close(c)
}
