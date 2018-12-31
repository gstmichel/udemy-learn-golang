package main

import "fmt"

type person struct {

}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I say...")
}

func saySomething(h human){
	h.speak()
}

func main(){
	person := person{}

	saySomething(&person)
}