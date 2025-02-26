package main

import "fmt"

/* call  by value 
func multiply(a,b int)int{

    return a * b
}
result:= multiply(5,20)
fmt.Println(result)

*/


// call by reference



func multiply(a,b *int)int{
	*a=*a*5
	return *a**b
}




func main() {

x:=50
y:=5
result:= multiply(&x,&y)

fmt.Println(result)



}