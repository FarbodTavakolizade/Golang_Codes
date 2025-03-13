package main

import (
    "fmt"
)
var factval uint64=1
var i int =1
var n int 

func factorial(n int )uint64 {
	if n < 0 {
		fmt.Print("invalid input")
	}else{
		for i:=1; i<=n; i++ {
			factval *= uint64(i)
        }
	}
	return factval  // return the factorial value 120 for n=5 120=5*4*3*2*1  and 120=6*5*4*3*2*1  and so on..
}


func main() {
fmt.Print("enter a positive integer number: ")
    fmt.Scan(&n)
fmt.Println(factorial(n))

}


