package main

import "fmt"

type customErr struct {

}

func (c customErr) Error() string {
	return "There's an error I guess..."
}

func main() {
	var c customErr
	foo(c)
}

func foo(err error){
	fmt.Printf("got error %v\n", err)
}