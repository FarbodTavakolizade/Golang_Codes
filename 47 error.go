package main

import (
	"fmt"
	"errors"
)

func main() {
	/*for i :=0 ;i<10;i++{
		result :=30/i
		fmt.Println(result)
	}*/

	x:="salam"
	MyError:=errors.New("invalid option")
	if x != "goodbye" {
		fmt.Println(MyError)
	}
}

//New()
//Erorf()
