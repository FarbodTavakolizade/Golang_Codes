package main

import (
    "fmt"
)


//method with pointer reciever

type person struct {
    name string
}

func(p*person)ChangeName(newName string) {
p.name = newName
}

func main() {
	x:=person{name: "John"}
	fmt.Println(x.name)

	x.ChangeName("reza")
	fmt.Println(x.name)

}