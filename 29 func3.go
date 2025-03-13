package main

import (
    "fmt"
)

func helloworld(){
	fmt.Println("hello world")
}

func name(name string,age int) {
fmt.Println(name,age)

}

func add(a int,b int) int {
return a+b

}

func add2(x int,y int)(result int){
	result=x+y
    return result
}


func main() {


helloworld()
name("farbod",20)
fmt.Println(add(2,6))
total:=add2(1,4)
fmt.Println(total)

}