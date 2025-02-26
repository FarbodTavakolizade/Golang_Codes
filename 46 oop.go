package main

import (
    "fmt"
)

// use interface for demonstrating abstraction
type veichle interface {
    start()
    stop()
}
//encapsulation
type Car struct {
    brand string
    year int
}

//constructor method for car
func NewCar(brand string, year int) *Car {
    return &Car{brand:brand, year:year}
}

func (c Car) start() {
    fmt.Println(c.brand,"car is starting")
}

func (c Car) stop() {
    fmt.Println(c.brand,"car is stopping")
}
//encapsulation
type bike struct {
    brand string
    year int
}

//constructor method for bike
  func NewBIKE(brand string, year int) *bike {
    return &bike{brand:brand, year:year}
}

func (b bike) start() {
    fmt.Println(b.brand,"bike is starting")
}

func (b bike) stop() {
    fmt.Println(b.brand,"bike is stopping")
}

//polymorphism
func x(v veichle){
    v.start()
    v.stop()
}

// creating two instance/objects in main function for car and bike 

func main() {
	car:=NewCar("bwm",2015)
    bike:=NewBIKE("honda",2020)
    x(car)
    x(bike)
}