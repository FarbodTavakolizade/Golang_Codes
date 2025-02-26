package main

import (
	"fmt"
	
)

func main() {
//   slice_name := [ ] datatype {value}

//myslice := [] int {1,2,3}   //length =3    //capacity=3

/*myslice := [] int{}
fmt.Println(myslice)
fmt.Println(len(myslice))
fmt.Println(cap(myslice))
*/

/*myslice :=[] string{"my","name","is","farbod"}
fmt.Println(myslice)
fmt.Println(len(myslice))
fmt.Println(cap(myslice))
*/


//slice_name := array[start index:end index]

/*arr1:=[6]int{10,11,12,13,14,15}
myslice :=arr1[2:4]
fmt.Printf("myslice =%v\n",myslice)
fmt.Printf("len = %d\n",len(myslice))
fmt.Printf("cap = %d\n",cap(myslice))
*/


//slice_name := make ([] datatype ,len,cap)


slice :=make([]int,5,10)
fmt.Printf("slice = %v\n",slice)
fmt.Printf("cap = %d\n",cap(slice))
fmt.Printf("len = %d\n",len(slice))











}